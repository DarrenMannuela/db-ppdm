package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellCoreDescription(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_core_description")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_core_description

	for rows.Next() {
		var well_core_description dto.Well_core_description
		if err := rows.Scan(&well_core_description.Uwi, &well_core_description.Source, &well_core_description.Core_id, &well_core_description.Description_obs_no, &well_core_description.Active_ind, &well_core_description.Base_depth, &well_core_description.Base_depth_ouom, &well_core_description.Core_description_company, &well_core_description.Description_date, &well_core_description.Dip_angle, &well_core_description.Effective_date, &well_core_description.Expiry_date, &well_core_description.Interval_thickness, &well_core_description.Interval_thickness_ouom, &well_core_description.Lithology_desc, &well_core_description.Porosity_length, &well_core_description.Porosity_length_ouom, &well_core_description.Porosity_quality, &well_core_description.Porosity_type, &well_core_description.Ppdm_guid, &well_core_description.Remark, &well_core_description.Show_length, &well_core_description.Show_length_ouom, &well_core_description.Show_type, &well_core_description.Top_depth, &well_core_description.Top_depth_ouom, &well_core_description.Row_changed_by, &well_core_description.Row_changed_date, &well_core_description.Row_created_by, &well_core_description.Row_created_date, &well_core_description.Row_effective_date, &well_core_description.Row_expiry_date, &well_core_description.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_core_description to result
		result = append(result, well_core_description)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellCoreDescription(c *fiber.Ctx) error {
	var well_core_description dto.Well_core_description

	setDefaults(&well_core_description)

	if err := json.Unmarshal(c.Body(), &well_core_description); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_core_description values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33)")
	if err != nil {
		return err
	}
	well_core_description.Row_created_date = formatDate(well_core_description.Row_created_date)
	well_core_description.Row_changed_date = formatDate(well_core_description.Row_changed_date)
	well_core_description.Description_date = formatDateString(well_core_description.Description_date)
	well_core_description.Effective_date = formatDateString(well_core_description.Effective_date)
	well_core_description.Expiry_date = formatDateString(well_core_description.Expiry_date)
	well_core_description.Row_effective_date = formatDateString(well_core_description.Row_effective_date)
	well_core_description.Row_expiry_date = formatDateString(well_core_description.Row_expiry_date)






	rows, err := stmt.Exec(well_core_description.Uwi, well_core_description.Source, well_core_description.Core_id, well_core_description.Description_obs_no, well_core_description.Active_ind, well_core_description.Base_depth, well_core_description.Base_depth_ouom, well_core_description.Core_description_company, well_core_description.Description_date, well_core_description.Dip_angle, well_core_description.Effective_date, well_core_description.Expiry_date, well_core_description.Interval_thickness, well_core_description.Interval_thickness_ouom, well_core_description.Lithology_desc, well_core_description.Porosity_length, well_core_description.Porosity_length_ouom, well_core_description.Porosity_quality, well_core_description.Porosity_type, well_core_description.Ppdm_guid, well_core_description.Remark, well_core_description.Show_length, well_core_description.Show_length_ouom, well_core_description.Show_type, well_core_description.Top_depth, well_core_description.Top_depth_ouom, well_core_description.Row_changed_by, well_core_description.Row_changed_date, well_core_description.Row_created_by, well_core_description.Row_created_date, well_core_description.Row_effective_date, well_core_description.Row_expiry_date, well_core_description.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellCoreDescription(c *fiber.Ctx) error {
	var well_core_description dto.Well_core_description

	setDefaults(&well_core_description)

	if err := json.Unmarshal(c.Body(), &well_core_description); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_core_description.Uwi = id

    if well_core_description.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_core_description set source = :1, core_id = :2, description_obs_no = :3, active_ind = :4, base_depth = :5, base_depth_ouom = :6, core_description_company = :7, description_date = :8, dip_angle = :9, effective_date = :10, expiry_date = :11, interval_thickness = :12, interval_thickness_ouom = :13, lithology_desc = :14, porosity_length = :15, porosity_length_ouom = :16, porosity_quality = :17, porosity_type = :18, ppdm_guid = :19, remark = :20, show_length = :21, show_length_ouom = :22, show_type = :23, top_depth = :24, top_depth_ouom = :25, row_changed_by = :26, row_changed_date = :27, row_created_by = :28, row_effective_date = :29, row_expiry_date = :30, row_quality = :31 where uwi = :33")
	if err != nil {
		return err
	}

	well_core_description.Row_changed_date = formatDate(well_core_description.Row_changed_date)
	well_core_description.Description_date = formatDateString(well_core_description.Description_date)
	well_core_description.Effective_date = formatDateString(well_core_description.Effective_date)
	well_core_description.Expiry_date = formatDateString(well_core_description.Expiry_date)
	well_core_description.Row_effective_date = formatDateString(well_core_description.Row_effective_date)
	well_core_description.Row_expiry_date = formatDateString(well_core_description.Row_expiry_date)






	rows, err := stmt.Exec(well_core_description.Source, well_core_description.Core_id, well_core_description.Description_obs_no, well_core_description.Active_ind, well_core_description.Base_depth, well_core_description.Base_depth_ouom, well_core_description.Core_description_company, well_core_description.Description_date, well_core_description.Dip_angle, well_core_description.Effective_date, well_core_description.Expiry_date, well_core_description.Interval_thickness, well_core_description.Interval_thickness_ouom, well_core_description.Lithology_desc, well_core_description.Porosity_length, well_core_description.Porosity_length_ouom, well_core_description.Porosity_quality, well_core_description.Porosity_type, well_core_description.Ppdm_guid, well_core_description.Remark, well_core_description.Show_length, well_core_description.Show_length_ouom, well_core_description.Show_type, well_core_description.Top_depth, well_core_description.Top_depth_ouom, well_core_description.Row_changed_by, well_core_description.Row_changed_date, well_core_description.Row_created_by, well_core_description.Row_effective_date, well_core_description.Row_expiry_date, well_core_description.Row_quality, well_core_description.Uwi)
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

func PatchWellCoreDescription(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_core_description set "
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
		} else if key == "description_date" || key == "effective_date" || key == "expiry_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteWellCoreDescription(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_core_description dto.Well_core_description
	well_core_description.Uwi = id

	stmt, err := db.Prepare("delete from well_core_description where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_core_description.Uwi)
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


