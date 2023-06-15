package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSeisGroupComp(c *fiber.Ctx) error {
	rows, err := db.Query("select * from seis_group_comp")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Seis_group_comp

	for rows.Next() {
		var seis_group_comp dto.Seis_group_comp
		if err := rows.Scan(&seis_group_comp.Seis_group_set_subtype, &seis_group_comp.Seis_group_id, &seis_group_comp.Input_seis_set_subtype, &seis_group_comp.Input_seis_set_id, &seis_group_comp.Component_id, &seis_group_comp.Active_ind, &seis_group_comp.Corner1_lat, &seis_group_comp.Corner1_long, &seis_group_comp.Corner2_lat, &seis_group_comp.Corner2_long, &seis_group_comp.Corner3_lat, &seis_group_comp.Corner3_long, &seis_group_comp.Description, &seis_group_comp.Effective_date, &seis_group_comp.Exclusion_ind, &seis_group_comp.Expiry_date, &seis_group_comp.Inclusion_ind, &seis_group_comp.Ppdm_guid, &seis_group_comp.Remark, &seis_group_comp.Seis_group_type, &seis_group_comp.Source, &seis_group_comp.Row_changed_by, &seis_group_comp.Row_changed_date, &seis_group_comp.Row_created_by, &seis_group_comp.Row_created_date, &seis_group_comp.Row_effective_date, &seis_group_comp.Row_expiry_date, &seis_group_comp.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append seis_group_comp to result
		result = append(result, seis_group_comp)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSeisGroupComp(c *fiber.Ctx) error {
	var seis_group_comp dto.Seis_group_comp

	setDefaults(&seis_group_comp)

	if err := json.Unmarshal(c.Body(), &seis_group_comp); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into seis_group_comp values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28)")
	if err != nil {
		return err
	}
	seis_group_comp.Row_created_date = formatDate(seis_group_comp.Row_created_date)
	seis_group_comp.Row_changed_date = formatDate(seis_group_comp.Row_changed_date)
	seis_group_comp.Effective_date = formatDateString(seis_group_comp.Effective_date)
	seis_group_comp.Expiry_date = formatDateString(seis_group_comp.Expiry_date)
	seis_group_comp.Row_effective_date = formatDateString(seis_group_comp.Row_effective_date)
	seis_group_comp.Row_expiry_date = formatDateString(seis_group_comp.Row_expiry_date)






	rows, err := stmt.Exec(seis_group_comp.Seis_group_set_subtype, seis_group_comp.Seis_group_id, seis_group_comp.Input_seis_set_subtype, seis_group_comp.Input_seis_set_id, seis_group_comp.Component_id, seis_group_comp.Active_ind, seis_group_comp.Corner1_lat, seis_group_comp.Corner1_long, seis_group_comp.Corner2_lat, seis_group_comp.Corner2_long, seis_group_comp.Corner3_lat, seis_group_comp.Corner3_long, seis_group_comp.Description, seis_group_comp.Effective_date, seis_group_comp.Exclusion_ind, seis_group_comp.Expiry_date, seis_group_comp.Inclusion_ind, seis_group_comp.Ppdm_guid, seis_group_comp.Remark, seis_group_comp.Seis_group_type, seis_group_comp.Source, seis_group_comp.Row_changed_by, seis_group_comp.Row_changed_date, seis_group_comp.Row_created_by, seis_group_comp.Row_created_date, seis_group_comp.Row_effective_date, seis_group_comp.Row_expiry_date, seis_group_comp.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSeisGroupComp(c *fiber.Ctx) error {
	var seis_group_comp dto.Seis_group_comp

	setDefaults(&seis_group_comp)

	if err := json.Unmarshal(c.Body(), &seis_group_comp); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	seis_group_comp.Seis_group_set_subtype = id

    if seis_group_comp.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update seis_group_comp set seis_group_id = :1, input_seis_set_subtype = :2, input_seis_set_id = :3, component_id = :4, active_ind = :5, corner1_lat = :6, corner1_long = :7, corner2_lat = :8, corner2_long = :9, corner3_lat = :10, corner3_long = :11, description = :12, effective_date = :13, exclusion_ind = :14, expiry_date = :15, inclusion_ind = :16, ppdm_guid = :17, remark = :18, seis_group_type = :19, source = :20, row_changed_by = :21, row_changed_date = :22, row_created_by = :23, row_effective_date = :24, row_expiry_date = :25, row_quality = :26 where seis_group_set_subtype = :28")
	if err != nil {
		return err
	}

	seis_group_comp.Row_changed_date = formatDate(seis_group_comp.Row_changed_date)
	seis_group_comp.Effective_date = formatDateString(seis_group_comp.Effective_date)
	seis_group_comp.Expiry_date = formatDateString(seis_group_comp.Expiry_date)
	seis_group_comp.Row_effective_date = formatDateString(seis_group_comp.Row_effective_date)
	seis_group_comp.Row_expiry_date = formatDateString(seis_group_comp.Row_expiry_date)






	rows, err := stmt.Exec(seis_group_comp.Seis_group_id, seis_group_comp.Input_seis_set_subtype, seis_group_comp.Input_seis_set_id, seis_group_comp.Component_id, seis_group_comp.Active_ind, seis_group_comp.Corner1_lat, seis_group_comp.Corner1_long, seis_group_comp.Corner2_lat, seis_group_comp.Corner2_long, seis_group_comp.Corner3_lat, seis_group_comp.Corner3_long, seis_group_comp.Description, seis_group_comp.Effective_date, seis_group_comp.Exclusion_ind, seis_group_comp.Expiry_date, seis_group_comp.Inclusion_ind, seis_group_comp.Ppdm_guid, seis_group_comp.Remark, seis_group_comp.Seis_group_type, seis_group_comp.Source, seis_group_comp.Row_changed_by, seis_group_comp.Row_changed_date, seis_group_comp.Row_created_by, seis_group_comp.Row_effective_date, seis_group_comp.Row_expiry_date, seis_group_comp.Row_quality, seis_group_comp.Seis_group_set_subtype)
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

func PatchSeisGroupComp(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update seis_group_comp set "
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
	query += " where seis_group_set_subtype = :id"

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

func DeleteSeisGroupComp(c *fiber.Ctx) error {
	id := c.Params("id")
	var seis_group_comp dto.Seis_group_comp
	seis_group_comp.Seis_group_set_subtype = id

	stmt, err := db.Prepare("delete from seis_group_comp where seis_group_set_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(seis_group_comp.Seis_group_set_subtype)
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


