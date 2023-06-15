package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPoolVersionArea(c *fiber.Ctx) error {
	rows, err := db.Query("select * from pool_version_area")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Pool_version_area

	for rows.Next() {
		var pool_version_area dto.Pool_version_area
		if err := rows.Scan(&pool_version_area.Pool_id, &pool_version_area.Pool_source, &pool_version_area.Area_id, &pool_version_area.Area_type, &pool_version_area.Active_ind, &pool_version_area.Effective_date, &pool_version_area.Expiry_date, &pool_version_area.Ppdm_guid, &pool_version_area.Remark, &pool_version_area.Source, &pool_version_area.Row_changed_by, &pool_version_area.Row_changed_date, &pool_version_area.Row_created_by, &pool_version_area.Row_created_date, &pool_version_area.Row_effective_date, &pool_version_area.Row_expiry_date, &pool_version_area.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append pool_version_area to result
		result = append(result, pool_version_area)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPoolVersionArea(c *fiber.Ctx) error {
	var pool_version_area dto.Pool_version_area

	setDefaults(&pool_version_area)

	if err := json.Unmarshal(c.Body(), &pool_version_area); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into pool_version_area values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17)")
	if err != nil {
		return err
	}
	pool_version_area.Row_created_date = formatDate(pool_version_area.Row_created_date)
	pool_version_area.Row_changed_date = formatDate(pool_version_area.Row_changed_date)
	pool_version_area.Effective_date = formatDateString(pool_version_area.Effective_date)
	pool_version_area.Expiry_date = formatDateString(pool_version_area.Expiry_date)
	pool_version_area.Row_effective_date = formatDateString(pool_version_area.Row_effective_date)
	pool_version_area.Row_expiry_date = formatDateString(pool_version_area.Row_expiry_date)






	rows, err := stmt.Exec(pool_version_area.Pool_id, pool_version_area.Pool_source, pool_version_area.Area_id, pool_version_area.Area_type, pool_version_area.Active_ind, pool_version_area.Effective_date, pool_version_area.Expiry_date, pool_version_area.Ppdm_guid, pool_version_area.Remark, pool_version_area.Source, pool_version_area.Row_changed_by, pool_version_area.Row_changed_date, pool_version_area.Row_created_by, pool_version_area.Row_created_date, pool_version_area.Row_effective_date, pool_version_area.Row_expiry_date, pool_version_area.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePoolVersionArea(c *fiber.Ctx) error {
	var pool_version_area dto.Pool_version_area

	setDefaults(&pool_version_area)

	if err := json.Unmarshal(c.Body(), &pool_version_area); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	pool_version_area.Pool_id = id

    if pool_version_area.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update pool_version_area set pool_source = :1, area_id = :2, area_type = :3, active_ind = :4, effective_date = :5, expiry_date = :6, ppdm_guid = :7, remark = :8, source = :9, row_changed_by = :10, row_changed_date = :11, row_created_by = :12, row_effective_date = :13, row_expiry_date = :14, row_quality = :15 where pool_id = :17")
	if err != nil {
		return err
	}

	pool_version_area.Row_changed_date = formatDate(pool_version_area.Row_changed_date)
	pool_version_area.Effective_date = formatDateString(pool_version_area.Effective_date)
	pool_version_area.Expiry_date = formatDateString(pool_version_area.Expiry_date)
	pool_version_area.Row_effective_date = formatDateString(pool_version_area.Row_effective_date)
	pool_version_area.Row_expiry_date = formatDateString(pool_version_area.Row_expiry_date)






	rows, err := stmt.Exec(pool_version_area.Pool_source, pool_version_area.Area_id, pool_version_area.Area_type, pool_version_area.Active_ind, pool_version_area.Effective_date, pool_version_area.Expiry_date, pool_version_area.Ppdm_guid, pool_version_area.Remark, pool_version_area.Source, pool_version_area.Row_changed_by, pool_version_area.Row_changed_date, pool_version_area.Row_created_by, pool_version_area.Row_effective_date, pool_version_area.Row_expiry_date, pool_version_area.Row_quality, pool_version_area.Pool_id)
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

func PatchPoolVersionArea(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update pool_version_area set "
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
	query += " where pool_id = :id"

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

func DeletePoolVersionArea(c *fiber.Ctx) error {
	id := c.Params("id")
	var pool_version_area dto.Pool_version_area
	pool_version_area.Pool_id = id

	stmt, err := db.Prepare("delete from pool_version_area where pool_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(pool_version_area.Pool_id)
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


