package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellHorizDrill(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_horiz_drill")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_horiz_drill

	for rows.Next() {
		var well_horiz_drill dto.Well_horiz_drill
		if err := rows.Scan(&well_horiz_drill.Uwi, &well_horiz_drill.Source, &well_horiz_drill.Active_ind, &well_horiz_drill.Buildup_radius_type, &well_horiz_drill.Buildup_rate_degree, &well_horiz_drill.Buildup_rate_length, &well_horiz_drill.Buildup_rate_length_ouom, &well_horiz_drill.Contractor, &well_horiz_drill.Effective_date, &well_horiz_drill.Expiry_date, &well_horiz_drill.Horiz_displacement, &well_horiz_drill.Horiz_displacement_ouom, &well_horiz_drill.Horiz_drilling_reason, &well_horiz_drill.Horiz_drilling_type, &well_horiz_drill.Horiz_strat_unit_id, &well_horiz_drill.Lateral_hole_id, &well_horiz_drill.Lateral_hole_length, &well_horiz_drill.Lateral_hole_length_ouom, &well_horiz_drill.Max_deviation_angle, &well_horiz_drill.Pay_length, &well_horiz_drill.Pay_length_ouom, &well_horiz_drill.Ppdm_guid, &well_horiz_drill.Rat_hole_depth, &well_horiz_drill.Rat_hole_depth_ouom, &well_horiz_drill.Rat_hole_depth_type, &well_horiz_drill.Remark, &well_horiz_drill.Reservoir, &well_horiz_drill.Strat_name_set_id, &well_horiz_drill.Wb_length_in_form, &well_horiz_drill.Wb_length_in_form_ouom, &well_horiz_drill.Row_changed_by, &well_horiz_drill.Row_changed_date, &well_horiz_drill.Row_created_by, &well_horiz_drill.Row_created_date, &well_horiz_drill.Row_effective_date, &well_horiz_drill.Row_expiry_date, &well_horiz_drill.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_horiz_drill to result
		result = append(result, well_horiz_drill)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellHorizDrill(c *fiber.Ctx) error {
	var well_horiz_drill dto.Well_horiz_drill

	setDefaults(&well_horiz_drill)

	if err := json.Unmarshal(c.Body(), &well_horiz_drill); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_horiz_drill values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37)")
	if err != nil {
		return err
	}
	well_horiz_drill.Row_created_date = formatDate(well_horiz_drill.Row_created_date)
	well_horiz_drill.Row_changed_date = formatDate(well_horiz_drill.Row_changed_date)
	well_horiz_drill.Effective_date = formatDateString(well_horiz_drill.Effective_date)
	well_horiz_drill.Expiry_date = formatDateString(well_horiz_drill.Expiry_date)
	well_horiz_drill.Row_effective_date = formatDateString(well_horiz_drill.Row_effective_date)
	well_horiz_drill.Row_expiry_date = formatDateString(well_horiz_drill.Row_expiry_date)






	rows, err := stmt.Exec(well_horiz_drill.Uwi, well_horiz_drill.Source, well_horiz_drill.Active_ind, well_horiz_drill.Buildup_radius_type, well_horiz_drill.Buildup_rate_degree, well_horiz_drill.Buildup_rate_length, well_horiz_drill.Buildup_rate_length_ouom, well_horiz_drill.Contractor, well_horiz_drill.Effective_date, well_horiz_drill.Expiry_date, well_horiz_drill.Horiz_displacement, well_horiz_drill.Horiz_displacement_ouom, well_horiz_drill.Horiz_drilling_reason, well_horiz_drill.Horiz_drilling_type, well_horiz_drill.Horiz_strat_unit_id, well_horiz_drill.Lateral_hole_id, well_horiz_drill.Lateral_hole_length, well_horiz_drill.Lateral_hole_length_ouom, well_horiz_drill.Max_deviation_angle, well_horiz_drill.Pay_length, well_horiz_drill.Pay_length_ouom, well_horiz_drill.Ppdm_guid, well_horiz_drill.Rat_hole_depth, well_horiz_drill.Rat_hole_depth_ouom, well_horiz_drill.Rat_hole_depth_type, well_horiz_drill.Remark, well_horiz_drill.Reservoir, well_horiz_drill.Strat_name_set_id, well_horiz_drill.Wb_length_in_form, well_horiz_drill.Wb_length_in_form_ouom, well_horiz_drill.Row_changed_by, well_horiz_drill.Row_changed_date, well_horiz_drill.Row_created_by, well_horiz_drill.Row_created_date, well_horiz_drill.Row_effective_date, well_horiz_drill.Row_expiry_date, well_horiz_drill.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellHorizDrill(c *fiber.Ctx) error {
	var well_horiz_drill dto.Well_horiz_drill

	setDefaults(&well_horiz_drill)

	if err := json.Unmarshal(c.Body(), &well_horiz_drill); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_horiz_drill.Uwi = id

    if well_horiz_drill.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_horiz_drill set source = :1, active_ind = :2, buildup_radius_type = :3, buildup_rate_degree = :4, buildup_rate_length = :5, buildup_rate_length_ouom = :6, contractor = :7, effective_date = :8, expiry_date = :9, horiz_displacement = :10, horiz_displacement_ouom = :11, horiz_drilling_reason = :12, horiz_drilling_type = :13, horiz_strat_unit_id = :14, lateral_hole_id = :15, lateral_hole_length = :16, lateral_hole_length_ouom = :17, max_deviation_angle = :18, pay_length = :19, pay_length_ouom = :20, ppdm_guid = :21, rat_hole_depth = :22, rat_hole_depth_ouom = :23, rat_hole_depth_type = :24, remark = :25, reservoir = :26, strat_name_set_id = :27, wb_length_in_form = :28, wb_length_in_form_ouom = :29, row_changed_by = :30, row_changed_date = :31, row_created_by = :32, row_effective_date = :33, row_expiry_date = :34, row_quality = :35 where uwi = :37")
	if err != nil {
		return err
	}

	well_horiz_drill.Row_changed_date = formatDate(well_horiz_drill.Row_changed_date)
	well_horiz_drill.Effective_date = formatDateString(well_horiz_drill.Effective_date)
	well_horiz_drill.Expiry_date = formatDateString(well_horiz_drill.Expiry_date)
	well_horiz_drill.Row_effective_date = formatDateString(well_horiz_drill.Row_effective_date)
	well_horiz_drill.Row_expiry_date = formatDateString(well_horiz_drill.Row_expiry_date)






	rows, err := stmt.Exec(well_horiz_drill.Source, well_horiz_drill.Active_ind, well_horiz_drill.Buildup_radius_type, well_horiz_drill.Buildup_rate_degree, well_horiz_drill.Buildup_rate_length, well_horiz_drill.Buildup_rate_length_ouom, well_horiz_drill.Contractor, well_horiz_drill.Effective_date, well_horiz_drill.Expiry_date, well_horiz_drill.Horiz_displacement, well_horiz_drill.Horiz_displacement_ouom, well_horiz_drill.Horiz_drilling_reason, well_horiz_drill.Horiz_drilling_type, well_horiz_drill.Horiz_strat_unit_id, well_horiz_drill.Lateral_hole_id, well_horiz_drill.Lateral_hole_length, well_horiz_drill.Lateral_hole_length_ouom, well_horiz_drill.Max_deviation_angle, well_horiz_drill.Pay_length, well_horiz_drill.Pay_length_ouom, well_horiz_drill.Ppdm_guid, well_horiz_drill.Rat_hole_depth, well_horiz_drill.Rat_hole_depth_ouom, well_horiz_drill.Rat_hole_depth_type, well_horiz_drill.Remark, well_horiz_drill.Reservoir, well_horiz_drill.Strat_name_set_id, well_horiz_drill.Wb_length_in_form, well_horiz_drill.Wb_length_in_form_ouom, well_horiz_drill.Row_changed_by, well_horiz_drill.Row_changed_date, well_horiz_drill.Row_created_by, well_horiz_drill.Row_effective_date, well_horiz_drill.Row_expiry_date, well_horiz_drill.Row_quality, well_horiz_drill.Uwi)
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

func PatchWellHorizDrill(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_horiz_drill set "
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

func DeleteWellHorizDrill(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_horiz_drill dto.Well_horiz_drill
	well_horiz_drill.Uwi = id

	stmt, err := db.Prepare("delete from well_horiz_drill where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_horiz_drill.Uwi)
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


