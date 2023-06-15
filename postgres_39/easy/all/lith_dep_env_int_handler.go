package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetLithDepEnvInt(c *fiber.Ctx) error {
	rows, err := db.Query("select * from lith_dep_env_int")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Lith_dep_env_int

	for rows.Next() {
		var lith_dep_env_int dto.Lith_dep_env_int
		if err := rows.Scan(&lith_dep_env_int.Lithology_log_id, &lith_dep_env_int.Lith_log_source, &lith_dep_env_int.Depth_obs_no, &lith_dep_env_int.Active_ind, &lith_dep_env_int.Base_depth, &lith_dep_env_int.Base_depth_ouom, &lith_dep_env_int.Depositional_environment, &lith_dep_env_int.Effective_date, &lith_dep_env_int.Expiry_date, &lith_dep_env_int.Ppdm_guid, &lith_dep_env_int.Remark, &lith_dep_env_int.Top_depth, &lith_dep_env_int.Top_depth_ouom, &lith_dep_env_int.Row_changed_by, &lith_dep_env_int.Row_changed_date, &lith_dep_env_int.Row_created_by, &lith_dep_env_int.Row_created_date, &lith_dep_env_int.Row_effective_date, &lith_dep_env_int.Row_expiry_date, &lith_dep_env_int.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append lith_dep_env_int to result
		result = append(result, lith_dep_env_int)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetLithDepEnvInt(c *fiber.Ctx) error {
	var lith_dep_env_int dto.Lith_dep_env_int

	setDefaults(&lith_dep_env_int)

	if err := json.Unmarshal(c.Body(), &lith_dep_env_int); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into lith_dep_env_int values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20)")
	if err != nil {
		return err
	}
	lith_dep_env_int.Row_created_date = formatDate(lith_dep_env_int.Row_created_date)
	lith_dep_env_int.Row_changed_date = formatDate(lith_dep_env_int.Row_changed_date)
	lith_dep_env_int.Effective_date = formatDateString(lith_dep_env_int.Effective_date)
	lith_dep_env_int.Expiry_date = formatDateString(lith_dep_env_int.Expiry_date)
	lith_dep_env_int.Row_effective_date = formatDateString(lith_dep_env_int.Row_effective_date)
	lith_dep_env_int.Row_expiry_date = formatDateString(lith_dep_env_int.Row_expiry_date)






	rows, err := stmt.Exec(lith_dep_env_int.Lithology_log_id, lith_dep_env_int.Lith_log_source, lith_dep_env_int.Depth_obs_no, lith_dep_env_int.Active_ind, lith_dep_env_int.Base_depth, lith_dep_env_int.Base_depth_ouom, lith_dep_env_int.Depositional_environment, lith_dep_env_int.Effective_date, lith_dep_env_int.Expiry_date, lith_dep_env_int.Ppdm_guid, lith_dep_env_int.Remark, lith_dep_env_int.Top_depth, lith_dep_env_int.Top_depth_ouom, lith_dep_env_int.Row_changed_by, lith_dep_env_int.Row_changed_date, lith_dep_env_int.Row_created_by, lith_dep_env_int.Row_created_date, lith_dep_env_int.Row_effective_date, lith_dep_env_int.Row_expiry_date, lith_dep_env_int.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateLithDepEnvInt(c *fiber.Ctx) error {
	var lith_dep_env_int dto.Lith_dep_env_int

	setDefaults(&lith_dep_env_int)

	if err := json.Unmarshal(c.Body(), &lith_dep_env_int); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	lith_dep_env_int.Lithology_log_id = id

    if lith_dep_env_int.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update lith_dep_env_int set lith_log_source = :1, depth_obs_no = :2, active_ind = :3, base_depth = :4, base_depth_ouom = :5, depositional_environment = :6, effective_date = :7, expiry_date = :8, ppdm_guid = :9, remark = :10, top_depth = :11, top_depth_ouom = :12, row_changed_by = :13, row_changed_date = :14, row_created_by = :15, row_effective_date = :16, row_expiry_date = :17, row_quality = :18 where lithology_log_id = :20")
	if err != nil {
		return err
	}

	lith_dep_env_int.Row_changed_date = formatDate(lith_dep_env_int.Row_changed_date)
	lith_dep_env_int.Effective_date = formatDateString(lith_dep_env_int.Effective_date)
	lith_dep_env_int.Expiry_date = formatDateString(lith_dep_env_int.Expiry_date)
	lith_dep_env_int.Row_effective_date = formatDateString(lith_dep_env_int.Row_effective_date)
	lith_dep_env_int.Row_expiry_date = formatDateString(lith_dep_env_int.Row_expiry_date)






	rows, err := stmt.Exec(lith_dep_env_int.Lith_log_source, lith_dep_env_int.Depth_obs_no, lith_dep_env_int.Active_ind, lith_dep_env_int.Base_depth, lith_dep_env_int.Base_depth_ouom, lith_dep_env_int.Depositional_environment, lith_dep_env_int.Effective_date, lith_dep_env_int.Expiry_date, lith_dep_env_int.Ppdm_guid, lith_dep_env_int.Remark, lith_dep_env_int.Top_depth, lith_dep_env_int.Top_depth_ouom, lith_dep_env_int.Row_changed_by, lith_dep_env_int.Row_changed_date, lith_dep_env_int.Row_created_by, lith_dep_env_int.Row_effective_date, lith_dep_env_int.Row_expiry_date, lith_dep_env_int.Row_quality, lith_dep_env_int.Lithology_log_id)
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

func PatchLithDepEnvInt(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update lith_dep_env_int set "
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
	query += " where lithology_log_id = :id"

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

func DeleteLithDepEnvInt(c *fiber.Ctx) error {
	id := c.Params("id")
	var lith_dep_env_int dto.Lith_dep_env_int
	lith_dep_env_int.Lithology_log_id = id

	stmt, err := db.Prepare("delete from lith_dep_env_int where lithology_log_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(lith_dep_env_int.Lithology_log_id)
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


