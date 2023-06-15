package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetProject(c *fiber.Ctx) error {
	rows, err := db.Query("select * from project")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Project

	for rows.Next() {
		var project dto.Project
		if err := rows.Scan(&project.Project_id, &project.Active_ind, &project.Complete_date, &project.Confidential_ind, &project.Confidential_release_date, &project.Description, &project.Effective_date, &project.Expiry_date, &project.Field_station_ind, &project.Land_right_ind, &project.Pden_ind, &project.Ppdm_guid, &project.Project_name, &project.Project_num, &project.Project_type, &project.Remark, &project.Seis_set_ind, &project.Source, &project.Start_date, &project.Strat_column_ind, &project.Strat_interpretation_ind, &project.Well_ind, &project.Row_changed_by, &project.Row_changed_date, &project.Row_created_by, &project.Row_created_date, &project.Row_effective_date, &project.Row_expiry_date, &project.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append project to result
		result = append(result, project)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetProject(c *fiber.Ctx) error {
	var project dto.Project

	setDefaults(&project)

	if err := json.Unmarshal(c.Body(), &project); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into project values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29)")
	if err != nil {
		return err
	}
	project.Row_created_date = formatDate(project.Row_created_date)
	project.Row_changed_date = formatDate(project.Row_changed_date)
	project.Complete_date = formatDateString(project.Complete_date)
	project.Confidential_release_date = formatDateString(project.Confidential_release_date)
	project.Effective_date = formatDateString(project.Effective_date)
	project.Expiry_date = formatDateString(project.Expiry_date)
	project.Start_date = formatDateString(project.Start_date)
	project.Row_effective_date = formatDateString(project.Row_effective_date)
	project.Row_expiry_date = formatDateString(project.Row_expiry_date)






	rows, err := stmt.Exec(project.Project_id, project.Active_ind, project.Complete_date, project.Confidential_ind, project.Confidential_release_date, project.Description, project.Effective_date, project.Expiry_date, project.Field_station_ind, project.Land_right_ind, project.Pden_ind, project.Ppdm_guid, project.Project_name, project.Project_num, project.Project_type, project.Remark, project.Seis_set_ind, project.Source, project.Start_date, project.Strat_column_ind, project.Strat_interpretation_ind, project.Well_ind, project.Row_changed_by, project.Row_changed_date, project.Row_created_by, project.Row_created_date, project.Row_effective_date, project.Row_expiry_date, project.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateProject(c *fiber.Ctx) error {
	var project dto.Project

	setDefaults(&project)

	if err := json.Unmarshal(c.Body(), &project); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	project.Project_id = id

    if project.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update project set active_ind = :1, complete_date = :2, confidential_ind = :3, confidential_release_date = :4, description = :5, effective_date = :6, expiry_date = :7, field_station_ind = :8, land_right_ind = :9, pden_ind = :10, ppdm_guid = :11, project_name = :12, project_num = :13, project_type = :14, remark = :15, seis_set_ind = :16, source = :17, start_date = :18, strat_column_ind = :19, strat_interpretation_ind = :20, well_ind = :21, row_changed_by = :22, row_changed_date = :23, row_created_by = :24, row_effective_date = :25, row_expiry_date = :26, row_quality = :27 where project_id = :29")
	if err != nil {
		return err
	}

	project.Row_changed_date = formatDate(project.Row_changed_date)
	project.Complete_date = formatDateString(project.Complete_date)
	project.Confidential_release_date = formatDateString(project.Confidential_release_date)
	project.Effective_date = formatDateString(project.Effective_date)
	project.Expiry_date = formatDateString(project.Expiry_date)
	project.Start_date = formatDateString(project.Start_date)
	project.Row_effective_date = formatDateString(project.Row_effective_date)
	project.Row_expiry_date = formatDateString(project.Row_expiry_date)






	rows, err := stmt.Exec(project.Active_ind, project.Complete_date, project.Confidential_ind, project.Confidential_release_date, project.Description, project.Effective_date, project.Expiry_date, project.Field_station_ind, project.Land_right_ind, project.Pden_ind, project.Ppdm_guid, project.Project_name, project.Project_num, project.Project_type, project.Remark, project.Seis_set_ind, project.Source, project.Start_date, project.Strat_column_ind, project.Strat_interpretation_ind, project.Well_ind, project.Row_changed_by, project.Row_changed_date, project.Row_created_by, project.Row_effective_date, project.Row_expiry_date, project.Row_quality, project.Project_id)
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

func PatchProject(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update project set "
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
		} else if key == "complete_date" || key == "confidential_release_date" || key == "effective_date" || key == "expiry_date" || key == "start_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where project_id = :id"

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

func DeleteProject(c *fiber.Ctx) error {
	id := c.Params("id")
	var project dto.Project
	project.Project_id = id

	stmt, err := db.Prepare("delete from project where project_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(project.Project_id)
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


