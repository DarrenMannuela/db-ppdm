package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPdenInArea(c *fiber.Ctx) error {
	rows, err := db.Query("select * from pden_in_area")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Pden_in_area

	for rows.Next() {
		var pden_in_area dto.Pden_in_area
		if err := rows.Scan(&pden_in_area.Pden_subtype, &pden_in_area.Pden_id, &pden_in_area.Pden_source, &pden_in_area.Area_id, &pden_in_area.Area_type, &pden_in_area.Active_ind, &pden_in_area.Effective_date, &pden_in_area.Expiry_date, &pden_in_area.Ppdm_guid, &pden_in_area.Remark, &pden_in_area.Source, &pden_in_area.Row_changed_by, &pden_in_area.Row_changed_date, &pden_in_area.Row_created_by, &pden_in_area.Row_created_date, &pden_in_area.Row_effective_date, &pden_in_area.Row_expiry_date, &pden_in_area.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append pden_in_area to result
		result = append(result, pden_in_area)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPdenInArea(c *fiber.Ctx) error {
	var pden_in_area dto.Pden_in_area

	setDefaults(&pden_in_area)

	if err := json.Unmarshal(c.Body(), &pden_in_area); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into pden_in_area values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18)")
	if err != nil {
		return err
	}
	pden_in_area.Row_created_date = formatDate(pden_in_area.Row_created_date)
	pden_in_area.Row_changed_date = formatDate(pden_in_area.Row_changed_date)
	pden_in_area.Effective_date = formatDateString(pden_in_area.Effective_date)
	pden_in_area.Expiry_date = formatDateString(pden_in_area.Expiry_date)
	pden_in_area.Row_effective_date = formatDateString(pden_in_area.Row_effective_date)
	pden_in_area.Row_expiry_date = formatDateString(pden_in_area.Row_expiry_date)






	rows, err := stmt.Exec(pden_in_area.Pden_subtype, pden_in_area.Pden_id, pden_in_area.Pden_source, pden_in_area.Area_id, pden_in_area.Area_type, pden_in_area.Active_ind, pden_in_area.Effective_date, pden_in_area.Expiry_date, pden_in_area.Ppdm_guid, pden_in_area.Remark, pden_in_area.Source, pden_in_area.Row_changed_by, pden_in_area.Row_changed_date, pden_in_area.Row_created_by, pden_in_area.Row_created_date, pden_in_area.Row_effective_date, pden_in_area.Row_expiry_date, pden_in_area.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePdenInArea(c *fiber.Ctx) error {
	var pden_in_area dto.Pden_in_area

	setDefaults(&pden_in_area)

	if err := json.Unmarshal(c.Body(), &pden_in_area); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	pden_in_area.Pden_subtype = id

    if pden_in_area.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update pden_in_area set pden_id = :1, pden_source = :2, area_id = :3, area_type = :4, active_ind = :5, effective_date = :6, expiry_date = :7, ppdm_guid = :8, remark = :9, source = :10, row_changed_by = :11, row_changed_date = :12, row_created_by = :13, row_effective_date = :14, row_expiry_date = :15, row_quality = :16 where pden_subtype = :18")
	if err != nil {
		return err
	}

	pden_in_area.Row_changed_date = formatDate(pden_in_area.Row_changed_date)
	pden_in_area.Effective_date = formatDateString(pden_in_area.Effective_date)
	pden_in_area.Expiry_date = formatDateString(pden_in_area.Expiry_date)
	pden_in_area.Row_effective_date = formatDateString(pden_in_area.Row_effective_date)
	pden_in_area.Row_expiry_date = formatDateString(pden_in_area.Row_expiry_date)






	rows, err := stmt.Exec(pden_in_area.Pden_id, pden_in_area.Pden_source, pden_in_area.Area_id, pden_in_area.Area_type, pden_in_area.Active_ind, pden_in_area.Effective_date, pden_in_area.Expiry_date, pden_in_area.Ppdm_guid, pden_in_area.Remark, pden_in_area.Source, pden_in_area.Row_changed_by, pden_in_area.Row_changed_date, pden_in_area.Row_created_by, pden_in_area.Row_effective_date, pden_in_area.Row_expiry_date, pden_in_area.Row_quality, pden_in_area.Pden_subtype)
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

func PatchPdenInArea(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update pden_in_area set "
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
	query += " where pden_subtype = :id"

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

func DeletePdenInArea(c *fiber.Ctx) error {
	id := c.Params("id")
	var pden_in_area dto.Pden_in_area
	pden_in_area.Pden_subtype = id

	stmt, err := db.Prepare("delete from pden_in_area where pden_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(pden_in_area.Pden_subtype)
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


