package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPoolArea(c *fiber.Ctx) error {
	rows, err := db.Query("select * from pool_area")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Pool_area

	for rows.Next() {
		var pool_area dto.Pool_area
		if err := rows.Scan(&pool_area.Pool_id, &pool_area.Area_id, &pool_area.Area_type, &pool_area.Active_ind, &pool_area.Effective_date, &pool_area.Expiry_date, &pool_area.Ppdm_guid, &pool_area.Remark, &pool_area.Source, &pool_area.Row_changed_by, &pool_area.Row_changed_date, &pool_area.Row_created_by, &pool_area.Row_created_date, &pool_area.Row_effective_date, &pool_area.Row_expiry_date, &pool_area.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append pool_area to result
		result = append(result, pool_area)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPoolArea(c *fiber.Ctx) error {
	var pool_area dto.Pool_area

	setDefaults(&pool_area)

	if err := json.Unmarshal(c.Body(), &pool_area); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into pool_area values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16)")
	if err != nil {
		return err
	}
	pool_area.Row_created_date = formatDate(pool_area.Row_created_date)
	pool_area.Row_changed_date = formatDate(pool_area.Row_changed_date)
	pool_area.Effective_date = formatDateString(pool_area.Effective_date)
	pool_area.Expiry_date = formatDateString(pool_area.Expiry_date)
	pool_area.Row_effective_date = formatDateString(pool_area.Row_effective_date)
	pool_area.Row_expiry_date = formatDateString(pool_area.Row_expiry_date)






	rows, err := stmt.Exec(pool_area.Pool_id, pool_area.Area_id, pool_area.Area_type, pool_area.Active_ind, pool_area.Effective_date, pool_area.Expiry_date, pool_area.Ppdm_guid, pool_area.Remark, pool_area.Source, pool_area.Row_changed_by, pool_area.Row_changed_date, pool_area.Row_created_by, pool_area.Row_created_date, pool_area.Row_effective_date, pool_area.Row_expiry_date, pool_area.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePoolArea(c *fiber.Ctx) error {
	var pool_area dto.Pool_area

	setDefaults(&pool_area)

	if err := json.Unmarshal(c.Body(), &pool_area); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	pool_area.Pool_id = id

    if pool_area.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update pool_area set area_id = :1, area_type = :2, active_ind = :3, effective_date = :4, expiry_date = :5, ppdm_guid = :6, remark = :7, source = :8, row_changed_by = :9, row_changed_date = :10, row_created_by = :11, row_effective_date = :12, row_expiry_date = :13, row_quality = :14 where pool_id = :16")
	if err != nil {
		return err
	}

	pool_area.Row_changed_date = formatDate(pool_area.Row_changed_date)
	pool_area.Effective_date = formatDateString(pool_area.Effective_date)
	pool_area.Expiry_date = formatDateString(pool_area.Expiry_date)
	pool_area.Row_effective_date = formatDateString(pool_area.Row_effective_date)
	pool_area.Row_expiry_date = formatDateString(pool_area.Row_expiry_date)






	rows, err := stmt.Exec(pool_area.Area_id, pool_area.Area_type, pool_area.Active_ind, pool_area.Effective_date, pool_area.Expiry_date, pool_area.Ppdm_guid, pool_area.Remark, pool_area.Source, pool_area.Row_changed_by, pool_area.Row_changed_date, pool_area.Row_created_by, pool_area.Row_effective_date, pool_area.Row_expiry_date, pool_area.Row_quality, pool_area.Pool_id)
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

func PatchPoolArea(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update pool_area set "
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

func DeletePoolArea(c *fiber.Ctx) error {
	id := c.Params("id")
	var pool_area dto.Pool_area
	pool_area.Pool_id = id

	stmt, err := db.Prepare("delete from pool_area where pool_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(pool_area.Pool_id)
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


