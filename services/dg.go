package services

import (
	"context"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"

	models "Modifa2/models"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4/pgxpool"
)

//
/*DB ... */
type DB struct {
}

//
func (db *DB) SaveOnDB(functionnamewithschema string, m interface{}) (models.AdminResponse, error) {
	userReg := models.AdminResponse{}
	u := ConVertInterface(functionnamewithschema, m)
	ctx := context.Background()
	db1, _ := pgxpool.Connect(ctx, os.Getenv("ProjectMain"))
	defer db1.Close()
	//
	err := pgxscan.Select(ctx, db1, &userReg, u)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			fmt.Println(pgErr.Message) // => syntax error at end of input
			fmt.Println(pgErr)         // => syntax error at end of input
			fmt.Println(pgErr.Code)    // => 42601
		}
	}
	return userReg, nil
}

func (db *DB) DeleteUser(functionnamewithschema string, m interface{}) (models.DeleteUserResponse, error) {
	userReg := models.DeleteUserResponse{}
	u := ConVertInterface(functionnamewithschema, m)
	ctx := context.Background()
	db1, _ := pgxpool.Connect(ctx, os.Getenv("ProjectMain"))
	defer db1.Close()
	//
	err := pgxscan.Select(ctx, db1, &userReg, u)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			fmt.Println(pgErr.Message) // => syntax error at end of input
			fmt.Println(pgErr)         // => syntax error at end of input
			fmt.Println(pgErr.Code)    // => 42601
		}
	}
	return userReg, nil
}

func (db *DB) GetUser(functionnamewithschema string, m interface{}) ([]models.GetUserResponse, error) {
	userReg := []models.GetUserResponse{}
	u := ConVertInterface(functionnamewithschema, m)
	ctx := context.Background()
	db1, _ := pgxpool.Connect(ctx, os.Getenv("ProjectMain"))
	defer db1.Close()
	//
	err := pgxscan.Select(ctx, db1, &userReg, u)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			fmt.Println(pgErr.Message) // => syntax error at end of input
			fmt.Println(pgErr)         // => syntax error at end of input
			fmt.Println(pgErr.Code)    // => 42601
		}
	}
	return userReg, nil
}

//Activate Account
func (db *DB) ActivateAccount(functionnamewithschema string, m interface{}) (models.ActivateResponse, error) {
	userReg := models.ActivateResponse{}
	u := ConVertInterface(functionnamewithschema, m)
	ctx := context.Background()
	db1, _ := pgxpool.Connect(ctx, os.Getenv("ProjectMain"))
	defer db1.Close()
	//
	err := pgxscan.Select(ctx, db1, &userReg, u)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			fmt.Println(pgErr.Message) // => syntax error at end of input
			fmt.Println(pgErr)         // => syntax error at end of input
			fmt.Println(pgErr.Code)    // => 42601
		}
	}
	return userReg, nil
}

//Delete Account
func (db *DB) DeleteAccount(functionnamewithschema string, m interface{}) (models.DeleteResponse, error) {
	userReg := models.DeleteResponse{}
	u := ConVertInterface(functionnamewithschema, m)
	ctx := context.Background()
	db1, _ := pgxpool.Connect(ctx, os.Getenv("ProjectMain"))
	defer db1.Close()
	//
	err := pgxscan.Select(ctx, db1, &userReg, u)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			fmt.Println(pgErr.Message) // => syntax error at end of input
			fmt.Println(pgErr)         // => syntax error at end of input
			fmt.Println(pgErr.Code)    // => 42601
		}
	}
	return userReg, nil
}

//Forgot Password
func (db *DB) ForgotPassword(functionnamewithschema string, m interface{}) (models.ForgotPasswordResponse, error) {
	userReg := models.ForgotPasswordResponse{}
	u := ConVertInterface(functionnamewithschema, m)
	ctx := context.Background()
	db1, _ := pgxpool.Connect(ctx, os.Getenv("ProjectMain"))
	defer db1.Close()
	//
	err := pgxscan.Select(ctx, db1, &userReg, u)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			fmt.Println(pgErr.Message) // => syntax error at end of input
			fmt.Println(pgErr)         // => syntax error at end of input
			fmt.Println(pgErr.Code)    // => 42601
		}
	}
	return userReg, nil
}

// PropertyResponse
func (db *DB) RegisterProperty(functionnamewithschema string, m interface{}) (models.PropertyResponse, error) {
	userReg := models.PropertyResponse{}
	u := ConVertInterface(functionnamewithschema, m)
	ctx := context.Background()
	db1, _ := pgxpool.Connect(ctx, os.Getenv("ProjectMain"))
	defer db1.Close()
	//
	err := pgxscan.Select(ctx, db1, &userReg, u)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			fmt.Println(pgErr.Message) // => syntax error at end of input
			fmt.Println(pgErr)         // => syntax error at end of input
			fmt.Println(pgErr.Code)    // => 42601
		}
	}
	return userReg, nil
}

//GetPropertyResponse
func (db *DB) GetProperty(functionnamewithschema string, m interface{}) ([]models.GetPropertyResponse, error) {
	userReg := []models.GetPropertyResponse{}
	u := ConVertInterface(functionnamewithschema, m)
	ctx := context.Background()
	db1, _ := pgxpool.Connect(ctx, os.Getenv("ProjectMain"))
	defer db1.Close()
	//
	err := pgxscan.Select(ctx, db1, &userReg, u)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			fmt.Println(pgErr.Message) // => syntax error at end of input
			fmt.Println(pgErr)         // => syntax error at end of input
			fmt.Println(pgErr.Code)    // => 42601
		}
	}
	return userReg, nil
}

//GetPropertyResponse
func (db *DB) GetAllAdmins(functionnamewithschema string, m interface{}) ([]models.GetAllAdminResponse, error) {
	userReg := []models.GetAllAdminResponse{}
	u := ConVertInterface(functionnamewithschema, m)
	ctx := context.Background()
	db1, _ := pgxpool.Connect(ctx, os.Getenv("ProjectMain"))
	defer db1.Close()
	//
	err := pgxscan.Select(ctx, db1, &userReg, u)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			fmt.Println(pgErr.Message) // => syntax error at end of input
			fmt.Println(pgErr)         // => syntax error at end of input
			fmt.Println(pgErr.Code)    // => 42601
		}
	}
	return userReg, nil
}

func (db *DB) ReturnAssocuations(functionnamewithschema string, m interface{}) ([]models.Associations, error) {
	User := []models.Associations{}
	u := ConVertInterface(functionnamewithschema, m)
	ctx := context.Background()
	db1, _ := pgxpool.Connect(ctx, os.Getenv("ProjectMain"))
	defer db1.Close()
	//
	err := pgxscan.Select(ctx, db1, &User, u)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			fmt.Println(pgErr.Message) // => syntax error at end of input
			fmt.Println(pgErr)         // => syntax error at end of input
			fmt.Println(pgErr.Code)    // => 42601
		}
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	}

	return User, nil
}

//Convert Interface and return Query string
func ConVertInterface(funcstr string, m interface{}) string {
	q := "select * from " + funcstr + "("

	if m != nil {

		v := reflect.ValueOf(m)
		typeOfS := v.Type()
		for i := 0; i < v.NumField(); i++ {

			switch typeOfS.Field(i).Type.Name() {
			case "int", "int16", "int32", "int64", "int8":
				str := v.Field(i).Interface().(int64)
				strInt64 := strconv.FormatInt(str, 10)
				q += strInt64 + ","
			case "float64":
				str := v.Field(i).Interface().(float64)
				s := fmt.Sprintf("%f", str)
				q += s + ","
			case "bool":
				q += "'" + strconv.FormatBool(v.Field(i).Interface().(bool)) + "',"
			default:
				if v.Field(i).Interface().(string) == "" {
					q += "null,"
				} else {
					q += "'" + v.Field(i).Interface().(string) + "',"
				}
			}
		}

		q = q[0 : len(q)-len(",")]
	}

	q += ")"

	return q
}
