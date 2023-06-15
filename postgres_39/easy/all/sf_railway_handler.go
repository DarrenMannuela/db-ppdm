package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSfRailway(c *fiber.Ctx) error {
	rows, err := db.Query("select * from sf_railway")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Sf_railway

	for rows.Next() {
		var sf_railway dto.Sf_railway
		if err := rows.Scan(&sf_railway.Sf_subtype, &sf_railway.Railway_id, &sf_railway.Active_ind, &sf_railway.Effective_date, &sf_railway.Expiry_date, &sf_railway.Ppdm_guid, &sf_railway.Rail_gauge, &sf_railway.Rail_gauge_ouom, &sf_railway.Remark, &sf_railway.Source, &sf_railway.Row_changed_by, &sf_railway.Row_changed_date, &sf_railway.Row_created_by, &sf_railway.Row_created_date, &sf_railway.Row_effective_date, &sf_railway.Row_expiry_date, &sf_railway.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append sf_railway to result
		result = append(result, sf_railway)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSfRailway(c *fiber.Ctx) error {
	var sf_railway dto.Sf_railway

	setDefaults(&sf_railway)

	if err := json.Unmarshal(c.Body(), &sf_railway); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into sf_railway values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17)")
	if err != nil {
		return err
	}
	sf_railway.Row_created_date = formatDate(sf_railway.Row_created_date)
	sf_railway.Row_changed_date = formatDate(sf_railway.Row_changed_date)
	sf_railway.Effective_date = formatDateString(sf_railway.Effective_date)
	sf_railway.Expiry_date = formatDateString(sf_railway.Expiry_date)
	sf_railway.Row_effective_date = formatDateString(sf_railway.Row_effective_date)
	sf_railway.Row_expiry_date = formatDateString(sf_railway.Row_expiry_date)






	rows, err := stmt.Exec(sf_railway.Sf_subtype, sf_railway.Railway_id, sf_railway.Active_ind, sf_railway.Effective_date, sf_railway.Expiry_date, sf_railway.Ppdm_guid, sf_railway.Rail_gauge, sf_railway.Rail_gauge_ouom, sf_railway.Remark, sf_railway.Source, sf_railway.Row_changed_by, sf_railway.Row_changed_date, sf_railway.Row_created_by, sf_railway.Row_created_date, sf_railway.Row_effective_date, sf_railway.Row_expiry_date, sf_railway.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSfRailway(c *fiber.Ctx) error {
	var sf_railway dto.Sf_railway

	setDefaults(&sf_railway)

	if err := json.Unmarshal(c.Body(), &sf_railway); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	sf_railway.Sf_subtype = id

    if sf_railway.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update sf_railway set railway_id = :1, active_ind = :2, effective_date = :3, expiry_date = :4, ppdm_guid = :5, rail_gauge = :6, rail_gauge_ouom = :7, remark = :8, source = :9, row_changed_by = :10, row_changed_date = :11, row_created_by = :12, row_effective_date = :13, row_expiry_date = :14, row_quality = :15 where sf_subtype = :17")
	if err != nil {
		return err
	}

	sf_railway.Row_changed_date = formatDate(sf_railway.Row_changed_date)
	sf_railway.Effective_date = formatDateString(sf_railway.Effective_date)
	sf_railway.Expiry_date = formatDateString(sf_railway.Expiry_date)
	sf_railway.Row_effective_date = formatDateString(sf_railway.Row_effective_date)
	sf_railway.Row_expiry_date = formatDateString(sf_railway.Row_expiry_date)






	rows, err := stmt.Exec(sf_railway.Railway_id, sf_railway.Active_ind, sf_railway.Effective_date, sf_railway.Expiry_date, sf_railway.Ppdm_guid, sf_railway.Rail_gauge, sf_railway.Rail_gauge_ouom, sf_railway.Remark, sf_railway.Source, sf_railway.Row_changed_by, sf_railway.Row_changed_date, sf_railway.Row_created_by, sf_railway.Row_effective_date, sf_railway.Row_expiry_date, sf_railway.Row_quality, sf_railway.Sf_subtype)
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

func PatchSfRailway(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update sf_railway set "
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
	query += " where sf_subtype = :id"

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

func DeleteSfRailway(c *fiber.Ctx) error {
	id := c.Params("id")
	var sf_railway dto.Sf_railway
	sf_railway.Sf_subtype = id

	stmt, err := db.Prepare("delete from sf_railway where sf_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(sf_railway.Sf_subtype)
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


