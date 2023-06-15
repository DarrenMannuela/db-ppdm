package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetProjectAlias(c *fiber.Ctx) error {
	rows, err := db.Query("select * from project_alias")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Project_alias

	for rows.Next() {
		var project_alias dto.Project_alias
		if err := rows.Scan(&project_alias.Project_id, &project_alias.Alias_id, &project_alias.Abbreviation, &project_alias.Active_ind, &project_alias.Alias_long_name, &project_alias.Alias_reason_type, &project_alias.Alias_short_name, &project_alias.Alias_type, &project_alias.Amended_date, &project_alias.Created_date, &project_alias.Effective_date, &project_alias.Expiry_date, &project_alias.Project_id, &project_alias.Original_ind, &project_alias.Owner_ba_id, &project_alias.Ppdm_guid, &project_alias.Preferred_ind, &project_alias.Reason_desc, &project_alias.Remark, &project_alias.Source, &project_alias.Source_document_id, &project_alias.Struckoff_date, &project_alias.Sw_application_id, &project_alias.Use_rule, &project_alias.Row_changed_by, &project_alias.Row_changed_date, &project_alias.Row_created_by, &project_alias.Row_created_date, &project_alias.Row_effective_date, &project_alias.Row_expiry_date, &project_alias.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append project_alias to result
		result = append(result, project_alias)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetProjectAlias(c *fiber.Ctx) error {
	var project_alias dto.Project_alias

	setDefaults(&project_alias)

	if err := json.Unmarshal(c.Body(), &project_alias); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into project_alias values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31)")
	if err != nil {
		return err
	}
	project_alias.Row_created_date = formatDate(project_alias.Row_created_date)
	project_alias.Row_changed_date = formatDate(project_alias.Row_changed_date)
	project_alias.Amended_date = formatDateString(project_alias.Amended_date)
	project_alias.Created_date = formatDateString(project_alias.Created_date)
	project_alias.Effective_date = formatDateString(project_alias.Effective_date)
	project_alias.Expiry_date = formatDateString(project_alias.Expiry_date)
	project_alias.Struckoff_date = formatDateString(project_alias.Struckoff_date)
	project_alias.Row_effective_date = formatDateString(project_alias.Row_effective_date)
	project_alias.Row_expiry_date = formatDateString(project_alias.Row_expiry_date)






	rows, err := stmt.Exec(project_alias.Project_id, project_alias.Alias_id, project_alias.Abbreviation, project_alias.Active_ind, project_alias.Alias_long_name, project_alias.Alias_reason_type, project_alias.Alias_short_name, project_alias.Alias_type, project_alias.Amended_date, project_alias.Created_date, project_alias.Effective_date, project_alias.Expiry_date, project_alias.Project_id, project_alias.Original_ind, project_alias.Owner_ba_id, project_alias.Ppdm_guid, project_alias.Preferred_ind, project_alias.Reason_desc, project_alias.Remark, project_alias.Source, project_alias.Source_document_id, project_alias.Struckoff_date, project_alias.Sw_application_id, project_alias.Use_rule, project_alias.Row_changed_by, project_alias.Row_changed_date, project_alias.Row_created_by, project_alias.Row_created_date, project_alias.Row_effective_date, project_alias.Row_expiry_date, project_alias.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateProjectAlias(c *fiber.Ctx) error {
	var project_alias dto.Project_alias

	setDefaults(&project_alias)

	if err := json.Unmarshal(c.Body(), &project_alias); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	project_alias.Project_id = id

    if project_alias.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update project_alias set alias_id = :1, abbreviation = :2, active_ind = :3, alias_long_name = :4, alias_reason_type = :5, alias_short_name = :6, alias_type = :7, amended_date = :8, created_date = :9, effective_date = :10, expiry_date = :11, project_id = :12, original_ind = :13, owner_ba_id = :14, ppdm_guid = :15, preferred_ind = :16, reason_desc = :17, remark = :18, source = :19, source_document_id = :20, struckoff_date = :21, sw_application_id = :22, use_rule = :23, row_changed_by = :24, row_changed_date = :25, row_created_by = :26, row_effective_date = :27, row_expiry_date = :28, row_quality = :29 where project_id = :31")
	if err != nil {
		return err
	}

	project_alias.Row_changed_date = formatDate(project_alias.Row_changed_date)
	project_alias.Amended_date = formatDateString(project_alias.Amended_date)
	project_alias.Created_date = formatDateString(project_alias.Created_date)
	project_alias.Effective_date = formatDateString(project_alias.Effective_date)
	project_alias.Expiry_date = formatDateString(project_alias.Expiry_date)
	project_alias.Struckoff_date = formatDateString(project_alias.Struckoff_date)
	project_alias.Row_effective_date = formatDateString(project_alias.Row_effective_date)
	project_alias.Row_expiry_date = formatDateString(project_alias.Row_expiry_date)






	rows, err := stmt.Exec(project_alias.Alias_id, project_alias.Abbreviation, project_alias.Active_ind, project_alias.Alias_long_name, project_alias.Alias_reason_type, project_alias.Alias_short_name, project_alias.Alias_type, project_alias.Amended_date, project_alias.Created_date, project_alias.Effective_date, project_alias.Expiry_date, project_alias.Project_id, project_alias.Original_ind, project_alias.Owner_ba_id, project_alias.Ppdm_guid, project_alias.Preferred_ind, project_alias.Reason_desc, project_alias.Remark, project_alias.Source, project_alias.Source_document_id, project_alias.Struckoff_date, project_alias.Sw_application_id, project_alias.Use_rule, project_alias.Row_changed_by, project_alias.Row_changed_date, project_alias.Row_created_by, project_alias.Row_effective_date, project_alias.Row_expiry_date, project_alias.Row_quality, project_alias.Project_id)
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

func PatchProjectAlias(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update project_alias set "
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
		} else if key == "amended_date" || key == "created_date" || key == "effective_date" || key == "expiry_date" || key == "struckoff_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteProjectAlias(c *fiber.Ctx) error {
	id := c.Params("id")
	var project_alias dto.Project_alias
	project_alias.Project_id = id

	stmt, err := db.Prepare("delete from project_alias where project_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(project_alias.Project_id)
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


