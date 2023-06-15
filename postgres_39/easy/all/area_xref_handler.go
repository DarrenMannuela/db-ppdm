package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetAreaXref(c *fiber.Ctx) error {
	rows, err := db.Query("select * from area_xref")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Area_xref

	for rows.Next() {
		var area_xref dto.Area_xref
		if err := rows.Scan(&area_xref.Area_id1, &area_xref.Area_type1, &area_xref.Area_id2, &area_xref.Area_type2, &area_xref.Area_xref_obs_no, &area_xref.Active_ind, &area_xref.Area_xref_type, &area_xref.Effective_date, &area_xref.Expiry_date, &area_xref.Ppdm_guid, &area_xref.Remark, &area_xref.Source, &area_xref.Row_changed_by, &area_xref.Row_changed_date, &area_xref.Row_created_by, &area_xref.Row_created_date, &area_xref.Row_effective_date, &area_xref.Row_expiry_date, &area_xref.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append area_xref to result
		result = append(result, area_xref)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetAreaXref(c *fiber.Ctx) error {
	var area_xref dto.Area_xref

	setDefaults(&area_xref)

	if err := json.Unmarshal(c.Body(), &area_xref); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into area_xref values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19)")
	if err != nil {
		return err
	}
	area_xref.Row_created_date = formatDate(area_xref.Row_created_date)
	area_xref.Row_changed_date = formatDate(area_xref.Row_changed_date)
	area_xref.Effective_date = formatDateString(area_xref.Effective_date)
	area_xref.Expiry_date = formatDateString(area_xref.Expiry_date)
	area_xref.Row_effective_date = formatDateString(area_xref.Row_effective_date)
	area_xref.Row_expiry_date = formatDateString(area_xref.Row_expiry_date)






	rows, err := stmt.Exec(area_xref.Area_id1, area_xref.Area_type1, area_xref.Area_id2, area_xref.Area_type2, area_xref.Area_xref_obs_no, area_xref.Active_ind, area_xref.Area_xref_type, area_xref.Effective_date, area_xref.Expiry_date, area_xref.Ppdm_guid, area_xref.Remark, area_xref.Source, area_xref.Row_changed_by, area_xref.Row_changed_date, area_xref.Row_created_by, area_xref.Row_created_date, area_xref.Row_effective_date, area_xref.Row_expiry_date, area_xref.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateAreaXref(c *fiber.Ctx) error {
	var area_xref dto.Area_xref

	setDefaults(&area_xref)

	if err := json.Unmarshal(c.Body(), &area_xref); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	area_xref.Area_id1 = id

    if area_xref.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update area_xref set area_type1 = :1, area_id2 = :2, area_type2 = :3, area_xref_obs_no = :4, active_ind = :5, area_xref_type = :6, effective_date = :7, expiry_date = :8, ppdm_guid = :9, remark = :10, source = :11, row_changed_by = :12, row_changed_date = :13, row_created_by = :14, row_effective_date = :15, row_expiry_date = :16, row_quality = :17 where area_id1 = :19")
	if err != nil {
		return err
	}

	area_xref.Row_changed_date = formatDate(area_xref.Row_changed_date)
	area_xref.Effective_date = formatDateString(area_xref.Effective_date)
	area_xref.Expiry_date = formatDateString(area_xref.Expiry_date)
	area_xref.Row_effective_date = formatDateString(area_xref.Row_effective_date)
	area_xref.Row_expiry_date = formatDateString(area_xref.Row_expiry_date)






	rows, err := stmt.Exec(area_xref.Area_type1, area_xref.Area_id2, area_xref.Area_type2, area_xref.Area_xref_obs_no, area_xref.Active_ind, area_xref.Area_xref_type, area_xref.Effective_date, area_xref.Expiry_date, area_xref.Ppdm_guid, area_xref.Remark, area_xref.Source, area_xref.Row_changed_by, area_xref.Row_changed_date, area_xref.Row_created_by, area_xref.Row_effective_date, area_xref.Row_expiry_date, area_xref.Row_quality, area_xref.Area_id1)
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

func PatchAreaXref(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update area_xref set "
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
	query += " where area_id1 = :id"

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

func DeleteAreaXref(c *fiber.Ctx) error {
	id := c.Params("id")
	var area_xref dto.Area_xref
	area_xref.Area_id1 = id

	stmt, err := db.Prepare("delete from area_xref where area_id1 = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(area_xref.Area_id1)
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


