package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetObligType(c *fiber.Ctx) error {
	rows, err := db.Query("select * from oblig_type")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Oblig_type

	for rows.Next() {
		var oblig_type dto.Oblig_type
		if err := rows.Scan(&oblig_type.Obligation_id, &oblig_type.Obligation_seq_no, &oblig_type.Land_oblig_type, &oblig_type.Active_ind, &oblig_type.Effective_date, &oblig_type.Expiry_date, &oblig_type.Land_oblig_category, &oblig_type.Ppdm_guid, &oblig_type.Remark, &oblig_type.Source, &oblig_type.Row_changed_by, &oblig_type.Row_changed_date, &oblig_type.Row_created_by, &oblig_type.Row_created_date, &oblig_type.Row_effective_date, &oblig_type.Row_expiry_date, &oblig_type.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append oblig_type to result
		result = append(result, oblig_type)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetObligType(c *fiber.Ctx) error {
	var oblig_type dto.Oblig_type

	setDefaults(&oblig_type)

	if err := json.Unmarshal(c.Body(), &oblig_type); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into oblig_type values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17)")
	if err != nil {
		return err
	}
	oblig_type.Row_created_date = formatDate(oblig_type.Row_created_date)
	oblig_type.Row_changed_date = formatDate(oblig_type.Row_changed_date)
	oblig_type.Effective_date = formatDateString(oblig_type.Effective_date)
	oblig_type.Expiry_date = formatDateString(oblig_type.Expiry_date)
	oblig_type.Row_effective_date = formatDateString(oblig_type.Row_effective_date)
	oblig_type.Row_expiry_date = formatDateString(oblig_type.Row_expiry_date)






	rows, err := stmt.Exec(oblig_type.Obligation_id, oblig_type.Obligation_seq_no, oblig_type.Land_oblig_type, oblig_type.Active_ind, oblig_type.Effective_date, oblig_type.Expiry_date, oblig_type.Land_oblig_category, oblig_type.Ppdm_guid, oblig_type.Remark, oblig_type.Source, oblig_type.Row_changed_by, oblig_type.Row_changed_date, oblig_type.Row_created_by, oblig_type.Row_created_date, oblig_type.Row_effective_date, oblig_type.Row_expiry_date, oblig_type.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateObligType(c *fiber.Ctx) error {
	var oblig_type dto.Oblig_type

	setDefaults(&oblig_type)

	if err := json.Unmarshal(c.Body(), &oblig_type); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	oblig_type.Obligation_id = id

    if oblig_type.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update oblig_type set obligation_seq_no = :1, land_oblig_type = :2, active_ind = :3, effective_date = :4, expiry_date = :5, land_oblig_category = :6, ppdm_guid = :7, remark = :8, source = :9, row_changed_by = :10, row_changed_date = :11, row_created_by = :12, row_effective_date = :13, row_expiry_date = :14, row_quality = :15 where obligation_id = :17")
	if err != nil {
		return err
	}

	oblig_type.Row_changed_date = formatDate(oblig_type.Row_changed_date)
	oblig_type.Effective_date = formatDateString(oblig_type.Effective_date)
	oblig_type.Expiry_date = formatDateString(oblig_type.Expiry_date)
	oblig_type.Row_effective_date = formatDateString(oblig_type.Row_effective_date)
	oblig_type.Row_expiry_date = formatDateString(oblig_type.Row_expiry_date)






	rows, err := stmt.Exec(oblig_type.Obligation_seq_no, oblig_type.Land_oblig_type, oblig_type.Active_ind, oblig_type.Effective_date, oblig_type.Expiry_date, oblig_type.Land_oblig_category, oblig_type.Ppdm_guid, oblig_type.Remark, oblig_type.Source, oblig_type.Row_changed_by, oblig_type.Row_changed_date, oblig_type.Row_created_by, oblig_type.Row_effective_date, oblig_type.Row_expiry_date, oblig_type.Row_quality, oblig_type.Obligation_id)
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

func PatchObligType(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update oblig_type set "
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
	query += " where obligation_id = :id"

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

func DeleteObligType(c *fiber.Ctx) error {
	id := c.Params("id")
	var oblig_type dto.Oblig_type
	oblig_type.Obligation_id = id

	stmt, err := db.Prepare("delete from oblig_type where obligation_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(oblig_type.Obligation_id)
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


