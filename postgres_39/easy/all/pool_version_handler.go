package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPoolVersion(c *fiber.Ctx) error {
	rows, err := db.Query("select * from pool_version")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Pool_version

	for rows.Next() {
		var pool_version dto.Pool_version
		if err := rows.Scan(&pool_version.Pool_id, &pool_version.Pool_source, &pool_version.Active_ind, &pool_version.Business_associate_id, &pool_version.Current_status_date, &pool_version.Discovery_date, &pool_version.Effective_date, &pool_version.Expiry_date, &pool_version.Field_id, &pool_version.Pool_code, &pool_version.Pool_name, &pool_version.Pool_name_abbreviation, &pool_version.Pool_status, &pool_version.Pool_type, &pool_version.Ppdm_guid, &pool_version.Remark, &pool_version.Strat_name_set_id, &pool_version.Strat_unit_id, &pool_version.Row_changed_by, &pool_version.Row_changed_date, &pool_version.Row_created_by, &pool_version.Row_created_date, &pool_version.Row_effective_date, &pool_version.Row_expiry_date, &pool_version.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append pool_version to result
		result = append(result, pool_version)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPoolVersion(c *fiber.Ctx) error {
	var pool_version dto.Pool_version

	setDefaults(&pool_version)

	if err := json.Unmarshal(c.Body(), &pool_version); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into pool_version values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25)")
	if err != nil {
		return err
	}
	pool_version.Row_created_date = formatDate(pool_version.Row_created_date)
	pool_version.Row_changed_date = formatDate(pool_version.Row_changed_date)
	pool_version.Current_status_date = formatDateString(pool_version.Current_status_date)
	pool_version.Discovery_date = formatDateString(pool_version.Discovery_date)
	pool_version.Effective_date = formatDateString(pool_version.Effective_date)
	pool_version.Expiry_date = formatDateString(pool_version.Expiry_date)
	pool_version.Row_effective_date = formatDateString(pool_version.Row_effective_date)
	pool_version.Row_expiry_date = formatDateString(pool_version.Row_expiry_date)






	rows, err := stmt.Exec(pool_version.Pool_id, pool_version.Pool_source, pool_version.Active_ind, pool_version.Business_associate_id, pool_version.Current_status_date, pool_version.Discovery_date, pool_version.Effective_date, pool_version.Expiry_date, pool_version.Field_id, pool_version.Pool_code, pool_version.Pool_name, pool_version.Pool_name_abbreviation, pool_version.Pool_status, pool_version.Pool_type, pool_version.Ppdm_guid, pool_version.Remark, pool_version.Strat_name_set_id, pool_version.Strat_unit_id, pool_version.Row_changed_by, pool_version.Row_changed_date, pool_version.Row_created_by, pool_version.Row_created_date, pool_version.Row_effective_date, pool_version.Row_expiry_date, pool_version.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePoolVersion(c *fiber.Ctx) error {
	var pool_version dto.Pool_version

	setDefaults(&pool_version)

	if err := json.Unmarshal(c.Body(), &pool_version); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	pool_version.Pool_id = id

    if pool_version.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update pool_version set pool_source = :1, active_ind = :2, business_associate_id = :3, current_status_date = :4, discovery_date = :5, effective_date = :6, expiry_date = :7, field_id = :8, pool_code = :9, pool_name = :10, pool_name_abbreviation = :11, pool_status = :12, pool_type = :13, ppdm_guid = :14, remark = :15, strat_name_set_id = :16, strat_unit_id = :17, row_changed_by = :18, row_changed_date = :19, row_created_by = :20, row_effective_date = :21, row_expiry_date = :22, row_quality = :23 where pool_id = :25")
	if err != nil {
		return err
	}

	pool_version.Row_changed_date = formatDate(pool_version.Row_changed_date)
	pool_version.Current_status_date = formatDateString(pool_version.Current_status_date)
	pool_version.Discovery_date = formatDateString(pool_version.Discovery_date)
	pool_version.Effective_date = formatDateString(pool_version.Effective_date)
	pool_version.Expiry_date = formatDateString(pool_version.Expiry_date)
	pool_version.Row_effective_date = formatDateString(pool_version.Row_effective_date)
	pool_version.Row_expiry_date = formatDateString(pool_version.Row_expiry_date)






	rows, err := stmt.Exec(pool_version.Pool_source, pool_version.Active_ind, pool_version.Business_associate_id, pool_version.Current_status_date, pool_version.Discovery_date, pool_version.Effective_date, pool_version.Expiry_date, pool_version.Field_id, pool_version.Pool_code, pool_version.Pool_name, pool_version.Pool_name_abbreviation, pool_version.Pool_status, pool_version.Pool_type, pool_version.Ppdm_guid, pool_version.Remark, pool_version.Strat_name_set_id, pool_version.Strat_unit_id, pool_version.Row_changed_by, pool_version.Row_changed_date, pool_version.Row_created_by, pool_version.Row_effective_date, pool_version.Row_expiry_date, pool_version.Row_quality, pool_version.Pool_id)
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

func PatchPoolVersion(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update pool_version set "
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

func DeletePoolVersion(c *fiber.Ctx) error {
	id := c.Params("id")
	var pool_version dto.Pool_version
	pool_version.Pool_id = id

	stmt, err := db.Prepare("delete from pool_version where pool_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(pool_version.Pool_id)
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


