package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellDirSrvyComposite(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_dir_srvy_composite")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_dir_srvy_composite

	for rows.Next() {
		var well_dir_srvy_composite dto.Well_dir_srvy_composite
		if err := rows.Scan(&well_dir_srvy_composite.Composite_uwi, &well_dir_srvy_composite.Composite_survey_id, &well_dir_srvy_composite.Composite_source, &well_dir_srvy_composite.Composite_obs_no, &well_dir_srvy_composite.Active_ind, &well_dir_srvy_composite.Depth_obs_no_from, &well_dir_srvy_composite.Depth_obs_no_to, &well_dir_srvy_composite.Effective_date, &well_dir_srvy_composite.Expiry_date, &well_dir_srvy_composite.Input_source, &well_dir_srvy_composite.Input_survey_id, &well_dir_srvy_composite.Input_uwi, &well_dir_srvy_composite.Ppdm_guid, &well_dir_srvy_composite.Remark, &well_dir_srvy_composite.Source, &well_dir_srvy_composite.Row_changed_by, &well_dir_srvy_composite.Row_changed_date, &well_dir_srvy_composite.Row_created_by, &well_dir_srvy_composite.Row_created_date, &well_dir_srvy_composite.Row_effective_date, &well_dir_srvy_composite.Row_expiry_date, &well_dir_srvy_composite.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_dir_srvy_composite to result
		result = append(result, well_dir_srvy_composite)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellDirSrvyComposite(c *fiber.Ctx) error {
	var well_dir_srvy_composite dto.Well_dir_srvy_composite

	setDefaults(&well_dir_srvy_composite)

	if err := json.Unmarshal(c.Body(), &well_dir_srvy_composite); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_dir_srvy_composite values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22)")
	if err != nil {
		return err
	}
	well_dir_srvy_composite.Row_created_date = formatDate(well_dir_srvy_composite.Row_created_date)
	well_dir_srvy_composite.Row_changed_date = formatDate(well_dir_srvy_composite.Row_changed_date)
	well_dir_srvy_composite.Effective_date = formatDateString(well_dir_srvy_composite.Effective_date)
	well_dir_srvy_composite.Expiry_date = formatDateString(well_dir_srvy_composite.Expiry_date)
	well_dir_srvy_composite.Row_effective_date = formatDateString(well_dir_srvy_composite.Row_effective_date)
	well_dir_srvy_composite.Row_expiry_date = formatDateString(well_dir_srvy_composite.Row_expiry_date)






	rows, err := stmt.Exec(well_dir_srvy_composite.Composite_uwi, well_dir_srvy_composite.Composite_survey_id, well_dir_srvy_composite.Composite_source, well_dir_srvy_composite.Composite_obs_no, well_dir_srvy_composite.Active_ind, well_dir_srvy_composite.Depth_obs_no_from, well_dir_srvy_composite.Depth_obs_no_to, well_dir_srvy_composite.Effective_date, well_dir_srvy_composite.Expiry_date, well_dir_srvy_composite.Input_source, well_dir_srvy_composite.Input_survey_id, well_dir_srvy_composite.Input_uwi, well_dir_srvy_composite.Ppdm_guid, well_dir_srvy_composite.Remark, well_dir_srvy_composite.Source, well_dir_srvy_composite.Row_changed_by, well_dir_srvy_composite.Row_changed_date, well_dir_srvy_composite.Row_created_by, well_dir_srvy_composite.Row_created_date, well_dir_srvy_composite.Row_effective_date, well_dir_srvy_composite.Row_expiry_date, well_dir_srvy_composite.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellDirSrvyComposite(c *fiber.Ctx) error {
	var well_dir_srvy_composite dto.Well_dir_srvy_composite

	setDefaults(&well_dir_srvy_composite)

	if err := json.Unmarshal(c.Body(), &well_dir_srvy_composite); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_dir_srvy_composite.Composite_uwi = id

    if well_dir_srvy_composite.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_dir_srvy_composite set composite_survey_id = :1, composite_source = :2, composite_obs_no = :3, active_ind = :4, depth_obs_no_from = :5, depth_obs_no_to = :6, effective_date = :7, expiry_date = :8, input_source = :9, input_survey_id = :10, input_uwi = :11, ppdm_guid = :12, remark = :13, source = :14, row_changed_by = :15, row_changed_date = :16, row_created_by = :17, row_effective_date = :18, row_expiry_date = :19, row_quality = :20 where composite_uwi = :22")
	if err != nil {
		return err
	}

	well_dir_srvy_composite.Row_changed_date = formatDate(well_dir_srvy_composite.Row_changed_date)
	well_dir_srvy_composite.Effective_date = formatDateString(well_dir_srvy_composite.Effective_date)
	well_dir_srvy_composite.Expiry_date = formatDateString(well_dir_srvy_composite.Expiry_date)
	well_dir_srvy_composite.Row_effective_date = formatDateString(well_dir_srvy_composite.Row_effective_date)
	well_dir_srvy_composite.Row_expiry_date = formatDateString(well_dir_srvy_composite.Row_expiry_date)






	rows, err := stmt.Exec(well_dir_srvy_composite.Composite_survey_id, well_dir_srvy_composite.Composite_source, well_dir_srvy_composite.Composite_obs_no, well_dir_srvy_composite.Active_ind, well_dir_srvy_composite.Depth_obs_no_from, well_dir_srvy_composite.Depth_obs_no_to, well_dir_srvy_composite.Effective_date, well_dir_srvy_composite.Expiry_date, well_dir_srvy_composite.Input_source, well_dir_srvy_composite.Input_survey_id, well_dir_srvy_composite.Input_uwi, well_dir_srvy_composite.Ppdm_guid, well_dir_srvy_composite.Remark, well_dir_srvy_composite.Source, well_dir_srvy_composite.Row_changed_by, well_dir_srvy_composite.Row_changed_date, well_dir_srvy_composite.Row_created_by, well_dir_srvy_composite.Row_effective_date, well_dir_srvy_composite.Row_expiry_date, well_dir_srvy_composite.Row_quality, well_dir_srvy_composite.Composite_uwi)
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

func PatchWellDirSrvyComposite(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_dir_srvy_composite set "
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
	query += " where composite_uwi = :id"

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

func DeleteWellDirSrvyComposite(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_dir_srvy_composite dto.Well_dir_srvy_composite
	well_dir_srvy_composite.Composite_uwi = id

	stmt, err := db.Prepare("delete from well_dir_srvy_composite where composite_uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_dir_srvy_composite.Composite_uwi)
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


