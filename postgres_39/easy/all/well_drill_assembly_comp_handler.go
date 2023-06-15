package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellDrillAssemblyComp(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_drill_assembly_comp")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_drill_assembly_comp

	for rows.Next() {
		var well_drill_assembly_comp dto.Well_drill_assembly_comp
		if err := rows.Scan(&well_drill_assembly_comp.Uwi, &well_drill_assembly_comp.Assembly_id, &well_drill_assembly_comp.Component_id, &well_drill_assembly_comp.Active_ind, &well_drill_assembly_comp.Catalogue_equip_id, &well_drill_assembly_comp.Component_count, &well_drill_assembly_comp.Component_desc, &well_drill_assembly_comp.Component_length, &well_drill_assembly_comp.Component_length_ouom, &well_drill_assembly_comp.Component_seq_no, &well_drill_assembly_comp.Component_type, &well_drill_assembly_comp.Component_weight, &well_drill_assembly_comp.Component_weight_ouom, &well_drill_assembly_comp.Effective_date, &well_drill_assembly_comp.Expiry_date, &well_drill_assembly_comp.Max_inside_diameter, &well_drill_assembly_comp.Max_inside_diameter_ouom, &well_drill_assembly_comp.Max_outside_diameter, &well_drill_assembly_comp.Max_outside_diameter_ouom, &well_drill_assembly_comp.Min_inside_diameter, &well_drill_assembly_comp.Min_inside_diameter_ouom, &well_drill_assembly_comp.Min_outside_diameter, &well_drill_assembly_comp.Min_outside_diameter_ouom, &well_drill_assembly_comp.Ppdm_guid, &well_drill_assembly_comp.Remark, &well_drill_assembly_comp.Serial_number, &well_drill_assembly_comp.Source, &well_drill_assembly_comp.Row_changed_by, &well_drill_assembly_comp.Row_changed_date, &well_drill_assembly_comp.Row_created_by, &well_drill_assembly_comp.Row_created_date, &well_drill_assembly_comp.Row_effective_date, &well_drill_assembly_comp.Row_expiry_date, &well_drill_assembly_comp.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_drill_assembly_comp to result
		result = append(result, well_drill_assembly_comp)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellDrillAssemblyComp(c *fiber.Ctx) error {
	var well_drill_assembly_comp dto.Well_drill_assembly_comp

	setDefaults(&well_drill_assembly_comp)

	if err := json.Unmarshal(c.Body(), &well_drill_assembly_comp); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_drill_assembly_comp values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34)")
	if err != nil {
		return err
	}
	well_drill_assembly_comp.Row_created_date = formatDate(well_drill_assembly_comp.Row_created_date)
	well_drill_assembly_comp.Row_changed_date = formatDate(well_drill_assembly_comp.Row_changed_date)
	well_drill_assembly_comp.Effective_date = formatDateString(well_drill_assembly_comp.Effective_date)
	well_drill_assembly_comp.Expiry_date = formatDateString(well_drill_assembly_comp.Expiry_date)
	well_drill_assembly_comp.Row_effective_date = formatDateString(well_drill_assembly_comp.Row_effective_date)
	well_drill_assembly_comp.Row_expiry_date = formatDateString(well_drill_assembly_comp.Row_expiry_date)






	rows, err := stmt.Exec(well_drill_assembly_comp.Uwi, well_drill_assembly_comp.Assembly_id, well_drill_assembly_comp.Component_id, well_drill_assembly_comp.Active_ind, well_drill_assembly_comp.Catalogue_equip_id, well_drill_assembly_comp.Component_count, well_drill_assembly_comp.Component_desc, well_drill_assembly_comp.Component_length, well_drill_assembly_comp.Component_length_ouom, well_drill_assembly_comp.Component_seq_no, well_drill_assembly_comp.Component_type, well_drill_assembly_comp.Component_weight, well_drill_assembly_comp.Component_weight_ouom, well_drill_assembly_comp.Effective_date, well_drill_assembly_comp.Expiry_date, well_drill_assembly_comp.Max_inside_diameter, well_drill_assembly_comp.Max_inside_diameter_ouom, well_drill_assembly_comp.Max_outside_diameter, well_drill_assembly_comp.Max_outside_diameter_ouom, well_drill_assembly_comp.Min_inside_diameter, well_drill_assembly_comp.Min_inside_diameter_ouom, well_drill_assembly_comp.Min_outside_diameter, well_drill_assembly_comp.Min_outside_diameter_ouom, well_drill_assembly_comp.Ppdm_guid, well_drill_assembly_comp.Remark, well_drill_assembly_comp.Serial_number, well_drill_assembly_comp.Source, well_drill_assembly_comp.Row_changed_by, well_drill_assembly_comp.Row_changed_date, well_drill_assembly_comp.Row_created_by, well_drill_assembly_comp.Row_created_date, well_drill_assembly_comp.Row_effective_date, well_drill_assembly_comp.Row_expiry_date, well_drill_assembly_comp.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellDrillAssemblyComp(c *fiber.Ctx) error {
	var well_drill_assembly_comp dto.Well_drill_assembly_comp

	setDefaults(&well_drill_assembly_comp)

	if err := json.Unmarshal(c.Body(), &well_drill_assembly_comp); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_drill_assembly_comp.Uwi = id

    if well_drill_assembly_comp.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_drill_assembly_comp set assembly_id = :1, component_id = :2, active_ind = :3, catalogue_equip_id = :4, component_count = :5, component_desc = :6, component_length = :7, component_length_ouom = :8, component_seq_no = :9, component_type = :10, component_weight = :11, component_weight_ouom = :12, effective_date = :13, expiry_date = :14, max_inside_diameter = :15, max_inside_diameter_ouom = :16, max_outside_diameter = :17, max_outside_diameter_ouom = :18, min_inside_diameter = :19, min_inside_diameter_ouom = :20, min_outside_diameter = :21, min_outside_diameter_ouom = :22, ppdm_guid = :23, remark = :24, serial_number = :25, source = :26, row_changed_by = :27, row_changed_date = :28, row_created_by = :29, row_effective_date = :30, row_expiry_date = :31, row_quality = :32 where uwi = :34")
	if err != nil {
		return err
	}

	well_drill_assembly_comp.Row_changed_date = formatDate(well_drill_assembly_comp.Row_changed_date)
	well_drill_assembly_comp.Effective_date = formatDateString(well_drill_assembly_comp.Effective_date)
	well_drill_assembly_comp.Expiry_date = formatDateString(well_drill_assembly_comp.Expiry_date)
	well_drill_assembly_comp.Row_effective_date = formatDateString(well_drill_assembly_comp.Row_effective_date)
	well_drill_assembly_comp.Row_expiry_date = formatDateString(well_drill_assembly_comp.Row_expiry_date)






	rows, err := stmt.Exec(well_drill_assembly_comp.Assembly_id, well_drill_assembly_comp.Component_id, well_drill_assembly_comp.Active_ind, well_drill_assembly_comp.Catalogue_equip_id, well_drill_assembly_comp.Component_count, well_drill_assembly_comp.Component_desc, well_drill_assembly_comp.Component_length, well_drill_assembly_comp.Component_length_ouom, well_drill_assembly_comp.Component_seq_no, well_drill_assembly_comp.Component_type, well_drill_assembly_comp.Component_weight, well_drill_assembly_comp.Component_weight_ouom, well_drill_assembly_comp.Effective_date, well_drill_assembly_comp.Expiry_date, well_drill_assembly_comp.Max_inside_diameter, well_drill_assembly_comp.Max_inside_diameter_ouom, well_drill_assembly_comp.Max_outside_diameter, well_drill_assembly_comp.Max_outside_diameter_ouom, well_drill_assembly_comp.Min_inside_diameter, well_drill_assembly_comp.Min_inside_diameter_ouom, well_drill_assembly_comp.Min_outside_diameter, well_drill_assembly_comp.Min_outside_diameter_ouom, well_drill_assembly_comp.Ppdm_guid, well_drill_assembly_comp.Remark, well_drill_assembly_comp.Serial_number, well_drill_assembly_comp.Source, well_drill_assembly_comp.Row_changed_by, well_drill_assembly_comp.Row_changed_date, well_drill_assembly_comp.Row_created_by, well_drill_assembly_comp.Row_effective_date, well_drill_assembly_comp.Row_expiry_date, well_drill_assembly_comp.Row_quality, well_drill_assembly_comp.Uwi)
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

func PatchWellDrillAssemblyComp(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_drill_assembly_comp set "
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

func DeleteWellDrillAssemblyComp(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_drill_assembly_comp dto.Well_drill_assembly_comp
	well_drill_assembly_comp.Uwi = id

	stmt, err := db.Prepare("delete from well_drill_assembly_comp where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_drill_assembly_comp.Uwi)
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


