package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellTestShutoff(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_test_shutoff")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_test_shutoff

	for rows.Next() {
		var well_test_shutoff dto.Well_test_shutoff
		if err := rows.Scan(&well_test_shutoff.Uwi, &well_test_shutoff.Source, &well_test_shutoff.Test_type, &well_test_shutoff.Run_num, &well_test_shutoff.Test_num, &well_test_shutoff.Shutoff_obs_no, &well_test_shutoff.Active_ind, &well_test_shutoff.Base_depth, &well_test_shutoff.Base_depth_ouom, &well_test_shutoff.Effective_date, &well_test_shutoff.Expiry_date, &well_test_shutoff.Plugback_depth, &well_test_shutoff.Plugback_depth_ouom, &well_test_shutoff.Ppdm_guid, &well_test_shutoff.Remark, &well_test_shutoff.Shutoff_type, &well_test_shutoff.Top_depth, &well_test_shutoff.Top_depth_ouom, &well_test_shutoff.Row_changed_by, &well_test_shutoff.Row_changed_date, &well_test_shutoff.Row_created_by, &well_test_shutoff.Row_created_date, &well_test_shutoff.Row_effective_date, &well_test_shutoff.Row_expiry_date, &well_test_shutoff.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_test_shutoff to result
		result = append(result, well_test_shutoff)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellTestShutoff(c *fiber.Ctx) error {
	var well_test_shutoff dto.Well_test_shutoff

	setDefaults(&well_test_shutoff)

	if err := json.Unmarshal(c.Body(), &well_test_shutoff); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_test_shutoff values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25)")
	if err != nil {
		return err
	}
	well_test_shutoff.Row_created_date = formatDate(well_test_shutoff.Row_created_date)
	well_test_shutoff.Row_changed_date = formatDate(well_test_shutoff.Row_changed_date)
	well_test_shutoff.Effective_date = formatDateString(well_test_shutoff.Effective_date)
	well_test_shutoff.Expiry_date = formatDateString(well_test_shutoff.Expiry_date)
	well_test_shutoff.Row_effective_date = formatDateString(well_test_shutoff.Row_effective_date)
	well_test_shutoff.Row_expiry_date = formatDateString(well_test_shutoff.Row_expiry_date)






	rows, err := stmt.Exec(well_test_shutoff.Uwi, well_test_shutoff.Source, well_test_shutoff.Test_type, well_test_shutoff.Run_num, well_test_shutoff.Test_num, well_test_shutoff.Shutoff_obs_no, well_test_shutoff.Active_ind, well_test_shutoff.Base_depth, well_test_shutoff.Base_depth_ouom, well_test_shutoff.Effective_date, well_test_shutoff.Expiry_date, well_test_shutoff.Plugback_depth, well_test_shutoff.Plugback_depth_ouom, well_test_shutoff.Ppdm_guid, well_test_shutoff.Remark, well_test_shutoff.Shutoff_type, well_test_shutoff.Top_depth, well_test_shutoff.Top_depth_ouom, well_test_shutoff.Row_changed_by, well_test_shutoff.Row_changed_date, well_test_shutoff.Row_created_by, well_test_shutoff.Row_created_date, well_test_shutoff.Row_effective_date, well_test_shutoff.Row_expiry_date, well_test_shutoff.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellTestShutoff(c *fiber.Ctx) error {
	var well_test_shutoff dto.Well_test_shutoff

	setDefaults(&well_test_shutoff)

	if err := json.Unmarshal(c.Body(), &well_test_shutoff); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_test_shutoff.Uwi = id

    if well_test_shutoff.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_test_shutoff set source = :1, test_type = :2, run_num = :3, test_num = :4, shutoff_obs_no = :5, active_ind = :6, base_depth = :7, base_depth_ouom = :8, effective_date = :9, expiry_date = :10, plugback_depth = :11, plugback_depth_ouom = :12, ppdm_guid = :13, remark = :14, shutoff_type = :15, top_depth = :16, top_depth_ouom = :17, row_changed_by = :18, row_changed_date = :19, row_created_by = :20, row_effective_date = :21, row_expiry_date = :22, row_quality = :23 where uwi = :25")
	if err != nil {
		return err
	}

	well_test_shutoff.Row_changed_date = formatDate(well_test_shutoff.Row_changed_date)
	well_test_shutoff.Effective_date = formatDateString(well_test_shutoff.Effective_date)
	well_test_shutoff.Expiry_date = formatDateString(well_test_shutoff.Expiry_date)
	well_test_shutoff.Row_effective_date = formatDateString(well_test_shutoff.Row_effective_date)
	well_test_shutoff.Row_expiry_date = formatDateString(well_test_shutoff.Row_expiry_date)






	rows, err := stmt.Exec(well_test_shutoff.Source, well_test_shutoff.Test_type, well_test_shutoff.Run_num, well_test_shutoff.Test_num, well_test_shutoff.Shutoff_obs_no, well_test_shutoff.Active_ind, well_test_shutoff.Base_depth, well_test_shutoff.Base_depth_ouom, well_test_shutoff.Effective_date, well_test_shutoff.Expiry_date, well_test_shutoff.Plugback_depth, well_test_shutoff.Plugback_depth_ouom, well_test_shutoff.Ppdm_guid, well_test_shutoff.Remark, well_test_shutoff.Shutoff_type, well_test_shutoff.Top_depth, well_test_shutoff.Top_depth_ouom, well_test_shutoff.Row_changed_by, well_test_shutoff.Row_changed_date, well_test_shutoff.Row_created_by, well_test_shutoff.Row_effective_date, well_test_shutoff.Row_expiry_date, well_test_shutoff.Row_quality, well_test_shutoff.Uwi)
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

func PatchWellTestShutoff(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_test_shutoff set "
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
	query += " where uwi = :id"

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

func DeleteWellTestShutoff(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_test_shutoff dto.Well_test_shutoff
	well_test_shutoff.Uwi = id

	stmt, err := db.Prepare("delete from well_test_shutoff where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_test_shutoff.Uwi)
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


