package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellDrillPipeInv(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_drill_pipe_inv")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_drill_pipe_inv

	for rows.Next() {
		var well_drill_pipe_inv dto.Well_drill_pipe_inv
		if err := rows.Scan(&well_drill_pipe_inv.Uwi, &well_drill_pipe_inv.Period_obs_no, &well_drill_pipe_inv.Inventory_obs_no, &well_drill_pipe_inv.Active_ind, &well_drill_pipe_inv.Cat_equip_id, &well_drill_pipe_inv.Coupling_type, &well_drill_pipe_inv.Damaged_joint_count, &well_drill_pipe_inv.Effective_date, &well_drill_pipe_inv.Expiry_date, &well_drill_pipe_inv.Joint_count, &well_drill_pipe_inv.Linear_mass, &well_drill_pipe_inv.Linear_mass_ouom, &well_drill_pipe_inv.Loan_authorized_by_ba_id, &well_drill_pipe_inv.Loan_count, &well_drill_pipe_inv.Loan_to_ba_id, &well_drill_pipe_inv.Lost_count, &well_drill_pipe_inv.Max_joint_outside_diam, &well_drill_pipe_inv.Max_joint_outside_diam_ouom, &well_drill_pipe_inv.Max_outside_diameter, &well_drill_pipe_inv.Max_outside_diameter_ouom, &well_drill_pipe_inv.Min_inside_diameter, &well_drill_pipe_inv.Min_inside_diameter_ouom, &well_drill_pipe_inv.Ppdm_guid, &well_drill_pipe_inv.Remark, &well_drill_pipe_inv.Report_time, &well_drill_pipe_inv.Report_timezone, &well_drill_pipe_inv.Report_time_type, &well_drill_pipe_inv.Source, &well_drill_pipe_inv.Tubing_grade, &well_drill_pipe_inv.Tubing_type, &well_drill_pipe_inv.Row_changed_by, &well_drill_pipe_inv.Row_changed_date, &well_drill_pipe_inv.Row_created_by, &well_drill_pipe_inv.Row_created_date, &well_drill_pipe_inv.Row_effective_date, &well_drill_pipe_inv.Row_expiry_date, &well_drill_pipe_inv.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_drill_pipe_inv to result
		result = append(result, well_drill_pipe_inv)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellDrillPipeInv(c *fiber.Ctx) error {
	var well_drill_pipe_inv dto.Well_drill_pipe_inv

	setDefaults(&well_drill_pipe_inv)

	if err := json.Unmarshal(c.Body(), &well_drill_pipe_inv); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_drill_pipe_inv values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37)")
	if err != nil {
		return err
	}
	well_drill_pipe_inv.Row_created_date = formatDate(well_drill_pipe_inv.Row_created_date)
	well_drill_pipe_inv.Row_changed_date = formatDate(well_drill_pipe_inv.Row_changed_date)
	well_drill_pipe_inv.Effective_date = formatDateString(well_drill_pipe_inv.Effective_date)
	well_drill_pipe_inv.Expiry_date = formatDateString(well_drill_pipe_inv.Expiry_date)
	well_drill_pipe_inv.Report_time = formatDateString(well_drill_pipe_inv.Report_time)
	well_drill_pipe_inv.Row_effective_date = formatDateString(well_drill_pipe_inv.Row_effective_date)
	well_drill_pipe_inv.Row_expiry_date = formatDateString(well_drill_pipe_inv.Row_expiry_date)






	rows, err := stmt.Exec(well_drill_pipe_inv.Uwi, well_drill_pipe_inv.Period_obs_no, well_drill_pipe_inv.Inventory_obs_no, well_drill_pipe_inv.Active_ind, well_drill_pipe_inv.Cat_equip_id, well_drill_pipe_inv.Coupling_type, well_drill_pipe_inv.Damaged_joint_count, well_drill_pipe_inv.Effective_date, well_drill_pipe_inv.Expiry_date, well_drill_pipe_inv.Joint_count, well_drill_pipe_inv.Linear_mass, well_drill_pipe_inv.Linear_mass_ouom, well_drill_pipe_inv.Loan_authorized_by_ba_id, well_drill_pipe_inv.Loan_count, well_drill_pipe_inv.Loan_to_ba_id, well_drill_pipe_inv.Lost_count, well_drill_pipe_inv.Max_joint_outside_diam, well_drill_pipe_inv.Max_joint_outside_diam_ouom, well_drill_pipe_inv.Max_outside_diameter, well_drill_pipe_inv.Max_outside_diameter_ouom, well_drill_pipe_inv.Min_inside_diameter, well_drill_pipe_inv.Min_inside_diameter_ouom, well_drill_pipe_inv.Ppdm_guid, well_drill_pipe_inv.Remark, well_drill_pipe_inv.Report_time, well_drill_pipe_inv.Report_timezone, well_drill_pipe_inv.Report_time_type, well_drill_pipe_inv.Source, well_drill_pipe_inv.Tubing_grade, well_drill_pipe_inv.Tubing_type, well_drill_pipe_inv.Row_changed_by, well_drill_pipe_inv.Row_changed_date, well_drill_pipe_inv.Row_created_by, well_drill_pipe_inv.Row_created_date, well_drill_pipe_inv.Row_effective_date, well_drill_pipe_inv.Row_expiry_date, well_drill_pipe_inv.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellDrillPipeInv(c *fiber.Ctx) error {
	var well_drill_pipe_inv dto.Well_drill_pipe_inv

	setDefaults(&well_drill_pipe_inv)

	if err := json.Unmarshal(c.Body(), &well_drill_pipe_inv); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_drill_pipe_inv.Uwi = id

    if well_drill_pipe_inv.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_drill_pipe_inv set period_obs_no = :1, inventory_obs_no = :2, active_ind = :3, cat_equip_id = :4, coupling_type = :5, damaged_joint_count = :6, effective_date = :7, expiry_date = :8, joint_count = :9, linear_mass = :10, linear_mass_ouom = :11, loan_authorized_by_ba_id = :12, loan_count = :13, loan_to_ba_id = :14, lost_count = :15, max_joint_outside_diam = :16, max_joint_outside_diam_ouom = :17, max_outside_diameter = :18, max_outside_diameter_ouom = :19, min_inside_diameter = :20, min_inside_diameter_ouom = :21, ppdm_guid = :22, remark = :23, report_time = :24, report_timezone = :25, report_time_type = :26, source = :27, tubing_grade = :28, tubing_type = :29, row_changed_by = :30, row_changed_date = :31, row_created_by = :32, row_effective_date = :33, row_expiry_date = :34, row_quality = :35 where uwi = :37")
	if err != nil {
		return err
	}

	well_drill_pipe_inv.Row_changed_date = formatDate(well_drill_pipe_inv.Row_changed_date)
	well_drill_pipe_inv.Effective_date = formatDateString(well_drill_pipe_inv.Effective_date)
	well_drill_pipe_inv.Expiry_date = formatDateString(well_drill_pipe_inv.Expiry_date)
	well_drill_pipe_inv.Report_time = formatDateString(well_drill_pipe_inv.Report_time)
	well_drill_pipe_inv.Row_effective_date = formatDateString(well_drill_pipe_inv.Row_effective_date)
	well_drill_pipe_inv.Row_expiry_date = formatDateString(well_drill_pipe_inv.Row_expiry_date)






	rows, err := stmt.Exec(well_drill_pipe_inv.Period_obs_no, well_drill_pipe_inv.Inventory_obs_no, well_drill_pipe_inv.Active_ind, well_drill_pipe_inv.Cat_equip_id, well_drill_pipe_inv.Coupling_type, well_drill_pipe_inv.Damaged_joint_count, well_drill_pipe_inv.Effective_date, well_drill_pipe_inv.Expiry_date, well_drill_pipe_inv.Joint_count, well_drill_pipe_inv.Linear_mass, well_drill_pipe_inv.Linear_mass_ouom, well_drill_pipe_inv.Loan_authorized_by_ba_id, well_drill_pipe_inv.Loan_count, well_drill_pipe_inv.Loan_to_ba_id, well_drill_pipe_inv.Lost_count, well_drill_pipe_inv.Max_joint_outside_diam, well_drill_pipe_inv.Max_joint_outside_diam_ouom, well_drill_pipe_inv.Max_outside_diameter, well_drill_pipe_inv.Max_outside_diameter_ouom, well_drill_pipe_inv.Min_inside_diameter, well_drill_pipe_inv.Min_inside_diameter_ouom, well_drill_pipe_inv.Ppdm_guid, well_drill_pipe_inv.Remark, well_drill_pipe_inv.Report_time, well_drill_pipe_inv.Report_timezone, well_drill_pipe_inv.Report_time_type, well_drill_pipe_inv.Source, well_drill_pipe_inv.Tubing_grade, well_drill_pipe_inv.Tubing_type, well_drill_pipe_inv.Row_changed_by, well_drill_pipe_inv.Row_changed_date, well_drill_pipe_inv.Row_created_by, well_drill_pipe_inv.Row_effective_date, well_drill_pipe_inv.Row_expiry_date, well_drill_pipe_inv.Row_quality, well_drill_pipe_inv.Uwi)
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

func PatchWellDrillPipeInv(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_drill_pipe_inv set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "report_time" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteWellDrillPipeInv(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_drill_pipe_inv dto.Well_drill_pipe_inv
	well_drill_pipe_inv.Uwi = id

	stmt, err := db.Prepare("delete from well_drill_pipe_inv where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_drill_pipe_inv.Uwi)
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


