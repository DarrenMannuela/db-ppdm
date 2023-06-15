package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetCsPrincipalMeridian(c *fiber.Ctx) error {
	rows, err := db.Query("select * from cs_principal_meridian")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Cs_principal_meridian

	for rows.Next() {
		var cs_principal_meridian dto.Cs_principal_meridian
		if err := rows.Scan(&cs_principal_meridian.Principal_meridian, &cs_principal_meridian.Abbreviation, &cs_principal_meridian.Active_ind, &cs_principal_meridian.Effective_date, &cs_principal_meridian.Expiry_date, &cs_principal_meridian.Longitude, &cs_principal_meridian.Long_name, &cs_principal_meridian.Ppdm_guid, &cs_principal_meridian.Remark, &cs_principal_meridian.Short_name, &cs_principal_meridian.Source, &cs_principal_meridian.Row_changed_by, &cs_principal_meridian.Row_changed_date, &cs_principal_meridian.Row_created_by, &cs_principal_meridian.Row_created_date, &cs_principal_meridian.Row_effective_date, &cs_principal_meridian.Row_expiry_date, &cs_principal_meridian.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append cs_principal_meridian to result
		result = append(result, cs_principal_meridian)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetCsPrincipalMeridian(c *fiber.Ctx) error {
	var cs_principal_meridian dto.Cs_principal_meridian

	setDefaults(&cs_principal_meridian)

	if err := json.Unmarshal(c.Body(), &cs_principal_meridian); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into cs_principal_meridian values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18)")
	if err != nil {
		return err
	}
	cs_principal_meridian.Row_created_date = formatDate(cs_principal_meridian.Row_created_date)
	cs_principal_meridian.Row_changed_date = formatDate(cs_principal_meridian.Row_changed_date)
	cs_principal_meridian.Effective_date = formatDateString(cs_principal_meridian.Effective_date)
	cs_principal_meridian.Expiry_date = formatDateString(cs_principal_meridian.Expiry_date)
	cs_principal_meridian.Row_effective_date = formatDateString(cs_principal_meridian.Row_effective_date)
	cs_principal_meridian.Row_expiry_date = formatDateString(cs_principal_meridian.Row_expiry_date)






	rows, err := stmt.Exec(cs_principal_meridian.Principal_meridian, cs_principal_meridian.Abbreviation, cs_principal_meridian.Active_ind, cs_principal_meridian.Effective_date, cs_principal_meridian.Expiry_date, cs_principal_meridian.Longitude, cs_principal_meridian.Long_name, cs_principal_meridian.Ppdm_guid, cs_principal_meridian.Remark, cs_principal_meridian.Short_name, cs_principal_meridian.Source, cs_principal_meridian.Row_changed_by, cs_principal_meridian.Row_changed_date, cs_principal_meridian.Row_created_by, cs_principal_meridian.Row_created_date, cs_principal_meridian.Row_effective_date, cs_principal_meridian.Row_expiry_date, cs_principal_meridian.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateCsPrincipalMeridian(c *fiber.Ctx) error {
	var cs_principal_meridian dto.Cs_principal_meridian

	setDefaults(&cs_principal_meridian)

	if err := json.Unmarshal(c.Body(), &cs_principal_meridian); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	cs_principal_meridian.Principal_meridian = id

    if cs_principal_meridian.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update cs_principal_meridian set abbreviation = :1, active_ind = :2, effective_date = :3, expiry_date = :4, longitude = :5, long_name = :6, ppdm_guid = :7, remark = :8, short_name = :9, source = :10, row_changed_by = :11, row_changed_date = :12, row_created_by = :13, row_effective_date = :14, row_expiry_date = :15, row_quality = :16 where principal_meridian = :18")
	if err != nil {
		return err
	}

	cs_principal_meridian.Row_changed_date = formatDate(cs_principal_meridian.Row_changed_date)
	cs_principal_meridian.Effective_date = formatDateString(cs_principal_meridian.Effective_date)
	cs_principal_meridian.Expiry_date = formatDateString(cs_principal_meridian.Expiry_date)
	cs_principal_meridian.Row_effective_date = formatDateString(cs_principal_meridian.Row_effective_date)
	cs_principal_meridian.Row_expiry_date = formatDateString(cs_principal_meridian.Row_expiry_date)






	rows, err := stmt.Exec(cs_principal_meridian.Abbreviation, cs_principal_meridian.Active_ind, cs_principal_meridian.Effective_date, cs_principal_meridian.Expiry_date, cs_principal_meridian.Longitude, cs_principal_meridian.Long_name, cs_principal_meridian.Ppdm_guid, cs_principal_meridian.Remark, cs_principal_meridian.Short_name, cs_principal_meridian.Source, cs_principal_meridian.Row_changed_by, cs_principal_meridian.Row_changed_date, cs_principal_meridian.Row_created_by, cs_principal_meridian.Row_effective_date, cs_principal_meridian.Row_expiry_date, cs_principal_meridian.Row_quality, cs_principal_meridian.Principal_meridian)
	if err != nil {
		return err
	}

	if count, err := rows.RowsAffected(); err == nil {
		if count == 0 {
			return c.Status(404).SendString("No matching record found")
		}
	} else {
		return err
	}

	return c.Status(201).JSON(rows)
}

func PatchCsPrincipalMeridian(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update cs_principal_meridian set "
	values := []interface{}{}
	i := 1
	for key, value := range updatedFields {
		query += key + " = :" + strconv.Itoa(i)
		i++
		if i <= len(updatedFields) {
			query += ", "
		}
		if key == "row_changed_date" {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDate(&date)
				value = formattedDate
			}
		} else if key == "effective_date" || key == "expiry_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where principal_meridian = :id"

	stmt, err := db.Prepare(query)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	values = append(values, id)
	res, err := stmt.Exec(values...)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	if count, err := res.RowsAffected(); err == nil {
		if count == 0 {
			return c.Status(404).SendString("No matching record found")
		}
	} else {
		return err
	}

	return c.JSON(res)
}

func DeleteCsPrincipalMeridian(c *fiber.Ctx) error {
	id := c.Params("id")
	var cs_principal_meridian dto.Cs_principal_meridian
	cs_principal_meridian.Principal_meridian = id

	stmt, err := db.Prepare("delete from cs_principal_meridian where principal_meridian = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(cs_principal_meridian.Principal_meridian)
	if err != nil {
		return err
	}

	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return c.Status(404).SendString("No matching record found")
	}

	return c.JSON(rows)
}


