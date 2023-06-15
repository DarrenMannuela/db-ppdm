package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPool(c *fiber.Ctx) error {
	rows, err := db.Query("select * from pool")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Pool

	for rows.Next() {
		var pool dto.Pool
		if err := rows.Scan(&pool.Pool_id, &pool.Active_ind, &pool.Business_associate_id, &pool.Current_status_date, &pool.Discovery_date, &pool.Effective_date, &pool.Expiry_date, &pool.Field_id, &pool.Pool_code, &pool.Pool_name, &pool.Pool_name_abbreviation, &pool.Pool_status, &pool.Pool_type, &pool.Ppdm_guid, &pool.Remark, &pool.Source, &pool.Strat_name_set_id, &pool.Strat_unit_id, &pool.Row_changed_by, &pool.Row_changed_date, &pool.Row_created_by, &pool.Row_created_date, &pool.Row_effective_date, &pool.Row_expiry_date, &pool.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append pool to result
		result = append(result, pool)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPool(c *fiber.Ctx) error {
	var pool dto.Pool

	setDefaults(&pool)

	if err := json.Unmarshal(c.Body(), &pool); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into pool values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25)")
	if err != nil {
		return err
	}
	pool.Row_created_date = formatDate(pool.Row_created_date)
	pool.Row_changed_date = formatDate(pool.Row_changed_date)
	pool.Current_status_date = formatDateString(pool.Current_status_date)
	pool.Discovery_date = formatDateString(pool.Discovery_date)
	pool.Effective_date = formatDateString(pool.Effective_date)
	pool.Expiry_date = formatDateString(pool.Expiry_date)
	pool.Row_effective_date = formatDateString(pool.Row_effective_date)
	pool.Row_expiry_date = formatDateString(pool.Row_expiry_date)






	rows, err := stmt.Exec(pool.Pool_id, pool.Active_ind, pool.Business_associate_id, pool.Current_status_date, pool.Discovery_date, pool.Effective_date, pool.Expiry_date, pool.Field_id, pool.Pool_code, pool.Pool_name, pool.Pool_name_abbreviation, pool.Pool_status, pool.Pool_type, pool.Ppdm_guid, pool.Remark, pool.Source, pool.Strat_name_set_id, pool.Strat_unit_id, pool.Row_changed_by, pool.Row_changed_date, pool.Row_created_by, pool.Row_created_date, pool.Row_effective_date, pool.Row_expiry_date, pool.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePool(c *fiber.Ctx) error {
	var pool dto.Pool

	setDefaults(&pool)

	if err := json.Unmarshal(c.Body(), &pool); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	pool.Pool_id = id

    if pool.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update pool set active_ind = :1, business_associate_id = :2, current_status_date = :3, discovery_date = :4, effective_date = :5, expiry_date = :6, field_id = :7, pool_code = :8, pool_name = :9, pool_name_abbreviation = :10, pool_status = :11, pool_type = :12, ppdm_guid = :13, remark = :14, source = :15, strat_name_set_id = :16, strat_unit_id = :17, row_changed_by = :18, row_changed_date = :19, row_created_by = :20, row_effective_date = :21, row_expiry_date = :22, row_quality = :23 where pool_id = :25")
	if err != nil {
		return err
	}

	pool.Row_changed_date = formatDate(pool.Row_changed_date)
	pool.Current_status_date = formatDateString(pool.Current_status_date)
	pool.Discovery_date = formatDateString(pool.Discovery_date)
	pool.Effective_date = formatDateString(pool.Effective_date)
	pool.Expiry_date = formatDateString(pool.Expiry_date)
	pool.Row_effective_date = formatDateString(pool.Row_effective_date)
	pool.Row_expiry_date = formatDateString(pool.Row_expiry_date)






	rows, err := stmt.Exec(pool.Active_ind, pool.Business_associate_id, pool.Current_status_date, pool.Discovery_date, pool.Effective_date, pool.Expiry_date, pool.Field_id, pool.Pool_code, pool.Pool_name, pool.Pool_name_abbreviation, pool.Pool_status, pool.Pool_type, pool.Ppdm_guid, pool.Remark, pool.Source, pool.Strat_name_set_id, pool.Strat_unit_id, pool.Row_changed_by, pool.Row_changed_date, pool.Row_created_by, pool.Row_effective_date, pool.Row_expiry_date, pool.Row_quality, pool.Pool_id)
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

func PatchPool(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update pool set "
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
		} else if key == "current_status_date" || key == "discovery_date" || key == "effective_date" || key == "expiry_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeletePool(c *fiber.Ctx) error {
	id := c.Params("id")
	var pool dto.Pool
	pool.Pool_id = id

	stmt, err := db.Prepare("delete from pool where pool_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(pool.Pool_id)
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


