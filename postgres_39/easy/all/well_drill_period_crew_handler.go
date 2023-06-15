package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellDrillPeriodCrew(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_drill_period_crew")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_drill_period_crew

	for rows.Next() {
		var well_drill_period_crew dto.Well_drill_period_crew
		if err := rows.Scan(&well_drill_period_crew.Uwi, &well_drill_period_crew.Period_obs_no, &well_drill_period_crew.Crew_company_ba_id, &well_drill_period_crew.Detail_obs_no, &well_drill_period_crew.Active_ind, &well_drill_period_crew.Booked_time_elapsed, &well_drill_period_crew.Booked_time_elapsed_uom, &well_drill_period_crew.Crew_member_ba_id, &well_drill_period_crew.Crew_member_name, &well_drill_period_crew.Crew_member_num, &well_drill_period_crew.Crew_member_record_ind, &well_drill_period_crew.Crew_position, &well_drill_period_crew.Crew_record_ind, &well_drill_period_crew.Crew_reference_num, &well_drill_period_crew.Effective_date, &well_drill_period_crew.Expiry_date, &well_drill_period_crew.Injury_ind, &well_drill_period_crew.Injury_initial_ind, &well_drill_period_crew.No_injury_ind, &well_drill_period_crew.Overhead_cost, &well_drill_period_crew.Overhead_cost_uom, &well_drill_period_crew.Overhead_type, &well_drill_period_crew.Ppdm_guid, &well_drill_period_crew.Remark, &well_drill_period_crew.Sf_subtype, &well_drill_period_crew.Source, &well_drill_period_crew.Subsistance_ind, &well_drill_period_crew.Support_facility_id, &well_drill_period_crew.Total_crew_count, &well_drill_period_crew.Total_injury_count, &well_drill_period_crew.Row_changed_by, &well_drill_period_crew.Row_changed_date, &well_drill_period_crew.Row_created_by, &well_drill_period_crew.Row_created_date, &well_drill_period_crew.Row_effective_date, &well_drill_period_crew.Row_expiry_date, &well_drill_period_crew.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_drill_period_crew to result
		result = append(result, well_drill_period_crew)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellDrillPeriodCrew(c *fiber.Ctx) error {
	var well_drill_period_crew dto.Well_drill_period_crew

	setDefaults(&well_drill_period_crew)

	if err := json.Unmarshal(c.Body(), &well_drill_period_crew); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_drill_period_crew values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37)")
	if err != nil {
		return err
	}
	well_drill_period_crew.Row_created_date = formatDate(well_drill_period_crew.Row_created_date)
	well_drill_period_crew.Row_changed_date = formatDate(well_drill_period_crew.Row_changed_date)
	well_drill_period_crew.Effective_date = formatDateString(well_drill_period_crew.Effective_date)
	well_drill_period_crew.Expiry_date = formatDateString(well_drill_period_crew.Expiry_date)
	well_drill_period_crew.Row_effective_date = formatDateString(well_drill_period_crew.Row_effective_date)
	well_drill_period_crew.Row_expiry_date = formatDateString(well_drill_period_crew.Row_expiry_date)






	rows, err := stmt.Exec(well_drill_period_crew.Uwi, well_drill_period_crew.Period_obs_no, well_drill_period_crew.Crew_company_ba_id, well_drill_period_crew.Detail_obs_no, well_drill_period_crew.Active_ind, well_drill_period_crew.Booked_time_elapsed, well_drill_period_crew.Booked_time_elapsed_uom, well_drill_period_crew.Crew_member_ba_id, well_drill_period_crew.Crew_member_name, well_drill_period_crew.Crew_member_num, well_drill_period_crew.Crew_member_record_ind, well_drill_period_crew.Crew_position, well_drill_period_crew.Crew_record_ind, well_drill_period_crew.Crew_reference_num, well_drill_period_crew.Effective_date, well_drill_period_crew.Expiry_date, well_drill_period_crew.Injury_ind, well_drill_period_crew.Injury_initial_ind, well_drill_period_crew.No_injury_ind, well_drill_period_crew.Overhead_cost, well_drill_period_crew.Overhead_cost_uom, well_drill_period_crew.Overhead_type, well_drill_period_crew.Ppdm_guid, well_drill_period_crew.Remark, well_drill_period_crew.Sf_subtype, well_drill_period_crew.Source, well_drill_period_crew.Subsistance_ind, well_drill_period_crew.Support_facility_id, well_drill_period_crew.Total_crew_count, well_drill_period_crew.Total_injury_count, well_drill_period_crew.Row_changed_by, well_drill_period_crew.Row_changed_date, well_drill_period_crew.Row_created_by, well_drill_period_crew.Row_created_date, well_drill_period_crew.Row_effective_date, well_drill_period_crew.Row_expiry_date, well_drill_period_crew.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellDrillPeriodCrew(c *fiber.Ctx) error {
	var well_drill_period_crew dto.Well_drill_period_crew

	setDefaults(&well_drill_period_crew)

	if err := json.Unmarshal(c.Body(), &well_drill_period_crew); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_drill_period_crew.Uwi = id

    if well_drill_period_crew.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_drill_period_crew set period_obs_no = :1, crew_company_ba_id = :2, detail_obs_no = :3, active_ind = :4, booked_time_elapsed = :5, booked_time_elapsed_uom = :6, crew_member_ba_id = :7, crew_member_name = :8, crew_member_num = :9, crew_member_record_ind = :10, crew_position = :11, crew_record_ind = :12, crew_reference_num = :13, effective_date = :14, expiry_date = :15, injury_ind = :16, injury_initial_ind = :17, no_injury_ind = :18, overhead_cost = :19, overhead_cost_uom = :20, overhead_type = :21, ppdm_guid = :22, remark = :23, sf_subtype = :24, source = :25, subsistance_ind = :26, support_facility_id = :27, total_crew_count = :28, total_injury_count = :29, row_changed_by = :30, row_changed_date = :31, row_created_by = :32, row_effective_date = :33, row_expiry_date = :34, row_quality = :35 where uwi = :37")
	if err != nil {
		return err
	}

	well_drill_period_crew.Row_changed_date = formatDate(well_drill_period_crew.Row_changed_date)
	well_drill_period_crew.Effective_date = formatDateString(well_drill_period_crew.Effective_date)
	well_drill_period_crew.Expiry_date = formatDateString(well_drill_period_crew.Expiry_date)
	well_drill_period_crew.Row_effective_date = formatDateString(well_drill_period_crew.Row_effective_date)
	well_drill_period_crew.Row_expiry_date = formatDateString(well_drill_period_crew.Row_expiry_date)






	rows, err := stmt.Exec(well_drill_period_crew.Period_obs_no, well_drill_period_crew.Crew_company_ba_id, well_drill_period_crew.Detail_obs_no, well_drill_period_crew.Active_ind, well_drill_period_crew.Booked_time_elapsed, well_drill_period_crew.Booked_time_elapsed_uom, well_drill_period_crew.Crew_member_ba_id, well_drill_period_crew.Crew_member_name, well_drill_period_crew.Crew_member_num, well_drill_period_crew.Crew_member_record_ind, well_drill_period_crew.Crew_position, well_drill_period_crew.Crew_record_ind, well_drill_period_crew.Crew_reference_num, well_drill_period_crew.Effective_date, well_drill_period_crew.Expiry_date, well_drill_period_crew.Injury_ind, well_drill_period_crew.Injury_initial_ind, well_drill_period_crew.No_injury_ind, well_drill_period_crew.Overhead_cost, well_drill_period_crew.Overhead_cost_uom, well_drill_period_crew.Overhead_type, well_drill_period_crew.Ppdm_guid, well_drill_period_crew.Remark, well_drill_period_crew.Sf_subtype, well_drill_period_crew.Source, well_drill_period_crew.Subsistance_ind, well_drill_period_crew.Support_facility_id, well_drill_period_crew.Total_crew_count, well_drill_period_crew.Total_injury_count, well_drill_period_crew.Row_changed_by, well_drill_period_crew.Row_changed_date, well_drill_period_crew.Row_created_by, well_drill_period_crew.Row_effective_date, well_drill_period_crew.Row_expiry_date, well_drill_period_crew.Row_quality, well_drill_period_crew.Uwi)
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

func PatchWellDrillPeriodCrew(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_drill_period_crew set "
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

func DeleteWellDrillPeriodCrew(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_drill_period_crew dto.Well_drill_period_crew
	well_drill_period_crew.Uwi = id

	stmt, err := db.Prepare("delete from well_drill_period_crew where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_drill_period_crew.Uwi)
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


