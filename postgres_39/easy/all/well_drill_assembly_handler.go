package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellDrillAssembly(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_drill_assembly")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_drill_assembly

	for rows.Next() {
		var well_drill_assembly dto.Well_drill_assembly
		if err := rows.Scan(&well_drill_assembly.Uwi, &well_drill_assembly.Assembly_id, &well_drill_assembly.Active_ind, &well_drill_assembly.Assembly_ref_number, &well_drill_assembly.Calculated_length, &well_drill_assembly.Calculated_weight, &well_drill_assembly.Effective_date, &well_drill_assembly.End_time, &well_drill_assembly.End_timezone, &well_drill_assembly.Expiry_date, &well_drill_assembly.Ppdm_guid, &well_drill_assembly.Remark, &well_drill_assembly.Source, &well_drill_assembly.Start_time, &well_drill_assembly.Start_timezone, &well_drill_assembly.Row_changed_by, &well_drill_assembly.Row_changed_date, &well_drill_assembly.Row_created_by, &well_drill_assembly.Row_created_date, &well_drill_assembly.Row_effective_date, &well_drill_assembly.Row_expiry_date, &well_drill_assembly.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_drill_assembly to result
		result = append(result, well_drill_assembly)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellDrillAssembly(c *fiber.Ctx) error {
	var well_drill_assembly dto.Well_drill_assembly

	setDefaults(&well_drill_assembly)

	if err := json.Unmarshal(c.Body(), &well_drill_assembly); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_drill_assembly values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22)")
	if err != nil {
		return err
	}
	well_drill_assembly.Row_created_date = formatDate(well_drill_assembly.Row_created_date)
	well_drill_assembly.Row_changed_date = formatDate(well_drill_assembly.Row_changed_date)
	well_drill_assembly.Effective_date = formatDateString(well_drill_assembly.Effective_date)
	well_drill_assembly.End_time = formatDateString(well_drill_assembly.End_time)
	well_drill_assembly.Expiry_date = formatDateString(well_drill_assembly.Expiry_date)
	well_drill_assembly.Start_time = formatDateString(well_drill_assembly.Start_time)
	well_drill_assembly.Row_effective_date = formatDateString(well_drill_assembly.Row_effective_date)
	well_drill_assembly.Row_expiry_date = formatDateString(well_drill_assembly.Row_expiry_date)






	rows, err := stmt.Exec(well_drill_assembly.Uwi, well_drill_assembly.Assembly_id, well_drill_assembly.Active_ind, well_drill_assembly.Assembly_ref_number, well_drill_assembly.Calculated_length, well_drill_assembly.Calculated_weight, well_drill_assembly.Effective_date, well_drill_assembly.End_time, well_drill_assembly.End_timezone, well_drill_assembly.Expiry_date, well_drill_assembly.Ppdm_guid, well_drill_assembly.Remark, well_drill_assembly.Source, well_drill_assembly.Start_time, well_drill_assembly.Start_timezone, well_drill_assembly.Row_changed_by, well_drill_assembly.Row_changed_date, well_drill_assembly.Row_created_by, well_drill_assembly.Row_created_date, well_drill_assembly.Row_effective_date, well_drill_assembly.Row_expiry_date, well_drill_assembly.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellDrillAssembly(c *fiber.Ctx) error {
	var well_drill_assembly dto.Well_drill_assembly

	setDefaults(&well_drill_assembly)

	if err := json.Unmarshal(c.Body(), &well_drill_assembly); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_drill_assembly.Uwi = id

    if well_drill_assembly.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_drill_assembly set assembly_id = :1, active_ind = :2, assembly_ref_number = :3, calculated_length = :4, calculated_weight = :5, effective_date = :6, end_time = :7, end_timezone = :8, expiry_date = :9, ppdm_guid = :10, remark = :11, source = :12, start_time = :13, start_timezone = :14, row_changed_by = :15, row_changed_date = :16, row_created_by = :17, row_effective_date = :18, row_expiry_date = :19, row_quality = :20 where uwi = :22")
	if err != nil {
		return err
	}

	well_drill_assembly.Row_changed_date = formatDate(well_drill_assembly.Row_changed_date)
	well_drill_assembly.Effective_date = formatDateString(well_drill_assembly.Effective_date)
	well_drill_assembly.End_time = formatDateString(well_drill_assembly.End_time)
	well_drill_assembly.Expiry_date = formatDateString(well_drill_assembly.Expiry_date)
	well_drill_assembly.Start_time = formatDateString(well_drill_assembly.Start_time)
	well_drill_assembly.Row_effective_date = formatDateString(well_drill_assembly.Row_effective_date)
	well_drill_assembly.Row_expiry_date = formatDateString(well_drill_assembly.Row_expiry_date)






	rows, err := stmt.Exec(well_drill_assembly.Assembly_id, well_drill_assembly.Active_ind, well_drill_assembly.Assembly_ref_number, well_drill_assembly.Calculated_length, well_drill_assembly.Calculated_weight, well_drill_assembly.Effective_date, well_drill_assembly.End_time, well_drill_assembly.End_timezone, well_drill_assembly.Expiry_date, well_drill_assembly.Ppdm_guid, well_drill_assembly.Remark, well_drill_assembly.Source, well_drill_assembly.Start_time, well_drill_assembly.Start_timezone, well_drill_assembly.Row_changed_by, well_drill_assembly.Row_changed_date, well_drill_assembly.Row_created_by, well_drill_assembly.Row_effective_date, well_drill_assembly.Row_expiry_date, well_drill_assembly.Row_quality, well_drill_assembly.Uwi)
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

func PatchWellDrillAssembly(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_drill_assembly set "
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
		} else if key == "effective_date" || key == "end_time" || key == "expiry_date" || key == "start_time" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteWellDrillAssembly(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_drill_assembly dto.Well_drill_assembly
	well_drill_assembly.Uwi = id

	stmt, err := db.Prepare("delete from well_drill_assembly where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_drill_assembly.Uwi)
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


