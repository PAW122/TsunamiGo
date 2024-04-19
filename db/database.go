package database

import (
	"database/sql"
	"errors"
	"math"
	types "tsunamiBot/types"

	_ "github.com/mattn/go-sqlite3"
)

func GetUserXp(data types.GetUserXp) types.GetUserXpRes {
	db, err := sql.Open("sqlite3", "./db/tsunami.db")
	if err != nil {
		return types.GetUserXpRes{Error: err}
	}
	defer db.Close()

	var res types.GetUserXpRes

	// Sprawdź dane użytkownika w bazie danych
	err = db.QueryRow("SELECT xp FROM Lvl WHERE user_id = ? AND guild_id = ?", data.UserId, data.GuildId).Scan(&res.Xp)
	if err != nil {
		if err == sql.ErrNoRows {
			// Jeśli użytkownik nie istnieje w bazie danych, zwróć błąd
			return types.GetUserXpRes{Error: errors.New("user not found")}
		}
		// Jeśli wystąpił inny błąd podczas zapytania do bazy danych, zwróć go
		return types.GetUserXpRes{Error: err}
	}

	// Oblicz poziom (Lvl) na podstawie punktów doświadczenia (XP)
	res.Lvl = CalculateLvl(res.Xp)
	res.Messages = res.Xp / 5
	res.UserId = data.UserId
	res.GuildId = data.GuildId

	return res
}

func CalculateLvl(xp int) int {
	// Tutaj możesz zaimplementować dowolny algorytm obliczania poziomu na podstawie punktów doświadczenia
	// Na przykład możesz użyć prostego algorytmu, który mówi, że każdy kolejny poziom jest osiągany po zdobyciu określonej liczby punktów doświadczenia
	// Na przykład, jeśli za każde 100 punktów użytkownik zdobywa nowy poziom, możesz użyć poniższego kodu:
	lvl := int(math.Floor(math.Sqrt(float64(xp) / 100)))
	return lvl
}

func SaveAddXp(data types.AddXp) error {
	// Otwórz lub stwórz nową bazę danych SQLite
	db, err := sql.Open("sqlite3", "./db/tsunami.db")
	if err != nil {
		return err
	}
	defer db.Close()

	user_id := data.UserId
	guild_id := data.GuildId

	// Utwórz tabelę Lvl jeśli nie istnieje
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS Lvl (
			user_id TEXT,
			guild_id TEXT,
			xp INTEGER,
			PRIMARY KEY (user_id, guild_id)
		)
	`)
	if err != nil {
		return err
	}

	// Sprawdź, czy użytkownik już istnieje w tabeli
	var currentXp int
	err = db.QueryRow("SELECT xp FROM Lvl WHERE user_id = ? AND guild_id = ?", user_id, guild_id).Scan(&currentXp)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	// Jeśli użytkownik już istnieje, dodaj nowe punkty XP do jego aktualnej liczby XP
	if err != sql.ErrNoRows {
		data.AddXp += currentXp
		_, err := db.Exec("UPDATE Lvl SET xp = ? WHERE user_id = ? AND guild_id = ?", data.AddXp, user_id, guild_id)
		if err != nil {
			return err
		}
		return nil
	}

	// Jeśli użytkownik nie istnieje, dodaj go do tabeli z nowymi punktami XP
	_, err = db.Exec("INSERT INTO Lvl (user_id, guild_id, xp) VALUES (?, ?, ?)", user_id, guild_id, data.AddXp)
	if err != nil {
		return err
	}

	return nil
}
