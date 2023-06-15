package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellDrillReport(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_drill_report")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_drill_report

	for rows.Next() {
		var well_drill_report dto.Well_drill_report
		if err := rows.Scan(&well_drill_report.Uwi, &well_drill_report.Report_id, &well_drill_report.Active_ind, &well_drill_report.Contractor_ba_id, &well_drill_report.Effective_date, &well_drill_report.End_date, &well_drill_report.End_time, &well_drill_report.End_timezone, &well_drill_report.Expiry_date, &well_drill_report.File_date, &well_drill_report.File_time, &well_drill_report.File_timezone, &well_drill_report.Period_count, &well_drill_report.Period_duration, &well_drill_report.Period_duration_uom, &well_drill_report.Ppdm_guid, &well_drill_report.Record_count, &well_drill_report.Reference_num, &well_drill_report.Remark, &well_drill_report.Reported_contractor_name, &well_drill_report.Reported_rig_num, &well_drill_report.Source, &well_drill_report.Standard_digital_format, &well_drill_report.Start_date, &well_drill_report.Start_time, &well_drill_report.Start_timezone, &well_drill_report.Sw_application_id, &well_drill_report.Vendor_digital_format, &well_drill_report.Version_seq_no, &well_drill_report.Row_changed_by, &well_drill_report.Row_changed_date, &well_drill_report.Row_created_by, &well_drill_report.Row_created_date, &well_drill_report.Row_effective_date, &well_drill_report.Row_expiry_date, &well_drill_report.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_drill_report to result
		result = append(result, well_drill_report)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellDrillReport(c *fiber.Ctx) error {
	var well_drill_report dto.Well_drill_report

	setDefaults(&well_drill_report)

	if err := json.Unmarshal(c.Body(), &well_drill_report); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_drill_report values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36)")
	if err != nil {
		return err
	}
	well_drill_report.Row_created_date = formatDate(well_drill_report.Row_created_date)
	well_drill_report.Row_changed_date = formatDate(well_drill_report.Row_changed_date)
	well_drill_report.Effective_date = formatDateString(well_drill_report.Effective_date)
	well_drill_report.End_date = formatDateString(well_drill_report.End_date)
	well_drill_report.End_time = formatDateString(well_drill_report.End_time)
	well_drill_report.Expiry_date = formatDateString(well_drill_report.Expiry_date)
	well_drill_report.File_date = formatDateString(well_drill_report.File_date)
	well_drill_report.File_time = formatDateString(well_drill_report.File_time)
	well_drill_report.Start_date = formatDateString(well_drill_report.Start_date)
	well_drill_report.Start_time = formatDateString(well_drill_report.Start_time)
	well_drill_report.Row_effective_date = formatDateString(well_drill_report.Row_effective_date)
	well_drill_report.Row_expiry_date = formatDateString(well_drill_report.Row_expiry_date)






	rows, err := stmt.Exec(well_drill_report.Uwi, well_drill_report.Report_id, well_drill_report.Active_ind, well_drill_report.Contractor_ba_id, well_drill_report.Effective_date, well_drill_report.End_date, well_drill_report.End_time, well_drill_report.End_timezone, well_drill_report.Expiry_date, well_drill_report.File_date, well_drill_report.File_time, well_drill_report.File_timezone, well_drill_report.Period_count, well_drill_report.Period_duration, well_drill_report.Period_duration_uom, well_drill_report.Ppdm_guid, well_drill_report.Record_count, well_drill_report.Reference_num, well_drill_report.Remark, well_drill_report.Reported_contractor_name, well_drill_report.Reported_rig_num, well_drill_report.Source, well_drill_report.Standard_digital_format, well_drill_report.Start_date, well_drill_report.Start_time, well_drill_report.Start_timezone, well_drill_report.Sw_application_id, well_drill_report.Vendor_digital_format, well_drill_report.Version_seq_no, well_drill_report.Row_changed_by, well_drill_report.Row_changed_date, well_drill_report.Row_created_by, well_drill_report.Row_created_date, well_drill_report.Row_effective_date, well_drill_report.Row_expiry_date, well_drill_report.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellDrillReport(c *fiber.Ctx) error {
	var well_drill_report dto.Well_drill_report

	setDefaults(&well_drill_report)

	if err := json.Unmarshal(c.Body(), &well_drill_report); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_drill_report.Uwi = id

    if well_drill_report.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_drill_report set report_id = :1, active_ind = :2, contractor_ba_id = :3, effective_date = :4, end_date = :5, end_time = :6, end_timezone = :7, expiry_date = :8, file_date = :9, file_time = :10, file_timezone = :11, period_count = :12, period_duration = :13, period_duration_uom = :14, ppdm_guid = :15, record_count = :16, reference_num = :17, remark = :18, reported_contractor_name = :19, reported_rig_num = :20, source = :21, standard_digital_format = :22, start_date = :23, start_time = :24, start_timezone = :25, sw_application_id = :26, vendor_digital_format = :27, version_seq_no = :28, row_changed_by = :29, row_changed_date = :30, row_created_by = :31, row_effective_date = :32, row_expiry_date = :33, row_quality = :34 where uwi = :36")
	if err != nil {
		return err
	}

	well_drill_report.Row_changed_date = formatDate(well_drill_report.Row_changed_date)
	well_drill_report.Effective_date = formatDateString(well_drill_report.Effective_date)
	well_drill_report.End_date = formatDateString(well_drill_report.End_date)
	well_drill_report.End_time = formatDateString(well_drill_report.End_time)
	well_drill_report.Expiry_date = formatDateString(well_drill_report.Expiry_date)
	well_drill_report.File_date = formatDateString(well_drill_report.File_date)
	well_drill_report.File_time = formatDateString(well_drill_report.File_time)
	well_drill_report.Start_date = formatDateString(well_drill_report.Start_date)
	well_drill_report.Start_time = formatDateString(well_drill_report.Start_time)
	well_drill_report.Row_effective_date = formatDateString(well_drill_report.Row_effective_date)
	well_drill_report.Row_expiry_date = formatDateString(well_drill_report.Row_expiry_date)






	rows, err := stmt.Exec(well_drill_report.Report_id, well_drill_report.Active_ind, well_drill_report.Contractor_ba_id, well_drill_report.Effective_date, well_drill_report.End_date, well_drill_report.End_time, well_drill_report.End_timezone, well_drill_report.Expiry_date, well_drill_report.File_date, well_drill_report.File_time, well_drill_report.File_timezone, well_drill_report.Period_count, well_drill_report.Period_duration, well_drill_report.Period_duration_uom, well_drill_report.Ppdm_guid, well_drill_report.Record_count, well_drill_report.Reference_num, well_drill_report.Remark, well_drill_report.Reported_contractor_name, well_drill_report.Reported_rig_num, well_drill_report.Source, well_drill_report.Standard_digital_format, well_drill_report.Start_date, well_drill_report.Start_time, well_drill_report.Start_timezone, well_drill_report.Sw_application_id, well_drill_report.Vendor_digital_format, well_drill_report.Version_seq_no, well_drill_report.Row_changed_by, well_drill_report.Row_changed_date, well_drill_report.Row_created_by, well_drill_report.Row_effective_date, well_drill_report.Row_expiry_date, well_drill_report.Row_quality, well_drill_report.Uwi)
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

func PatchWellDrillReport(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_drill_report set "
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
		} else if key == "effective_date" || key == "end_date" || key == "end_time" || key == "expiry_date" || key == "file_date" || key == "file_time" || key == "start_date" || key == "start_time" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteWellDrillReport(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_drill_report dto.Well_drill_report
	well_drill_report.Uwi = id

	stmt, err := db.Prepare("delete from well_drill_report where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_drill_report.Uwi)
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


