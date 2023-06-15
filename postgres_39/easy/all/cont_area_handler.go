package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetContArea(c *fiber.Ctx) error {
	rows, err := db.Query("select * from cont_area")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Cont_area

	for rows.Next() {
		var cont_area dto.Cont_area
		if err := rows.Scan(&cont_area.Contract_id, &cont_area.Area_id, &cont_area.Area_type, &cont_area.Active_ind, &cont_area.Effective_date, &cont_area.Expiry_date, &cont_area.Ppdm_guid, &cont_area.Remark, &cont_area.Source, &cont_area.Row_changed_by, &cont_area.Row_changed_date, &cont_area.Row_created_by, &cont_area.Row_created_date, &cont_area.Row_effective_date, &cont_area.Row_expiry_date, &cont_area.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append cont_area to result
		result = append(result, cont_area)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetContArea(c *fiber.Ctx) error {
	var cont_area dto.Cont_area

	setDefaults(&cont_area)

	if err := json.Unmarshal(c.Body(), &cont_area); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into cont_area values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16)")
	if err != nil {
		return err
	}
	cont_area.Row_created_date = formatDate(cont_area.Row_created_date)
	cont_area.Row_changed_date = formatDate(cont_area.Row_changed_date)
	cont_area.Effective_date = formatDateString(cont_area.Effective_date)
	cont_area.Expiry_date = formatDateString(cont_area.Expiry_date)
	cont_area.Row_effective_date = formatDateString(cont_area.Row_effective_date)
	cont_area.Row_expiry_date = formatDateString(cont_area.Row_expiry_date)






	rows, err := stmt.Exec(cont_area.Contract_id, cont_area.Area_id, cont_area.Area_type, cont_area.Active_ind, cont_area.Effective_date, cont_area.Expiry_date, cont_area.Ppdm_guid, cont_area.Remark, cont_area.Source, cont_area.Row_changed_by, cont_area.Row_changed_date, cont_area.Row_created_by, cont_area.Row_created_date, cont_area.Row_effective_date, cont_area.Row_expiry_date, cont_area.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateContArea(c *fiber.Ctx) error {
	var cont_area dto.Cont_area

	setDefaults(&cont_area)

	if err := json.Unmarshal(c.Body(), &cont_area); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	cont_area.Contract_id = id

    if cont_area.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update cont_area set area_id = :1, area_type = :2, active_ind = :3, effective_date = :4, expiry_date = :5, ppdm_guid = :6, remark = :7, source = :8, row_changed_by = :9, row_changed_date = :10, row_created_by = :11, row_effective_date = :12, row_expiry_date = :13, row_quality = :14 where contract_id = :16")
	if err != nil {
		return err
	}

	cont_area.Row_changed_date = formatDate(cont_area.Row_changed_date)
	cont_area.Effective_date = formatDateString(cont_area.Effective_date)
	cont_area.Expiry_date = formatDateString(cont_area.Expiry_date)
	cont_area.Row_effective_date = formatDateString(cont_area.Row_effective_date)
	cont_area.Row_expiry_date = formatDateString(cont_area.Row_expiry_date)






	rows, err := stmt.Exec(cont_area.Area_id, cont_area.Area_type, cont_area.Active_ind, cont_area.Effective_date, cont_area.Expiry_date, cont_area.Ppdm_guid, cont_area.Remark, cont_area.Source, cont_area.Row_changed_by, cont_area.Row_changed_date, cont_area.Row_created_by, cont_area.Row_effective_date, cont_area.Row_expiry_date, cont_area.Row_quality, cont_area.Contract_id)
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

func PatchContArea(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update cont_area set "
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
	query += " where contract_id = :id"

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

func DeleteContArea(c *fiber.Ctx) error {
	id := c.Params("id")
	var cont_area dto.Cont_area
	cont_area.Contract_id = id

	stmt, err := db.Prepare("delete from cont_area where contract_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(cont_area.Contract_id)
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


