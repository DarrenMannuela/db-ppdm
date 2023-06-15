package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellDrillIntDetail(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_drill_int_detail")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_drill_int_detail

	for rows.Next() {
		var well_drill_int_detail dto.Well_drill_int_detail
		if err := rows.Scan(&well_drill_int_detail.Uwi, &well_drill_int_detail.Bit_source, &well_drill_int_detail.Bit_interval_obs_no, &well_drill_int_detail.Bit_detail_type, &well_drill_int_detail.Detail_obs_no, &well_drill_int_detail.Active_ind, &well_drill_int_detail.Average_value, &well_drill_int_detail.Average_value_ouom, &well_drill_int_detail.Average_value_uom, &well_drill_int_detail.Base_depth, &well_drill_int_detail.Base_depth_ouom, &well_drill_int_detail.Bit_detail_code, &well_drill_int_detail.Date_format_desc, &well_drill_int_detail.Detail_desc, &well_drill_int_detail.Effective_date, &well_drill_int_detail.End_date, &well_drill_int_detail.End_time, &well_drill_int_detail.Expiry_date, &well_drill_int_detail.Max_value, &well_drill_int_detail.Max_value_ouom, &well_drill_int_detail.Max_value_uom, &well_drill_int_detail.Min_value, &well_drill_int_detail.Min_value_ouom, &well_drill_int_detail.Min_value_uom, &well_drill_int_detail.Ppdm_guid, &well_drill_int_detail.Remark, &well_drill_int_detail.Source, &well_drill_int_detail.Start_date, &well_drill_int_detail.Start_time, &well_drill_int_detail.Timezone, &well_drill_int_detail.Top_depth, &well_drill_int_detail.Top_depth_ouom, &well_drill_int_detail.Row_changed_by, &well_drill_int_detail.Row_changed_date, &well_drill_int_detail.Row_created_by, &well_drill_int_detail.Row_created_date, &well_drill_int_detail.Row_effective_date, &well_drill_int_detail.Row_expiry_date, &well_drill_int_detail.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_drill_int_detail to result
		result = append(result, well_drill_int_detail)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellDrillIntDetail(c *fiber.Ctx) error {
	var well_drill_int_detail dto.Well_drill_int_detail

	setDefaults(&well_drill_int_detail)

	if err := json.Unmarshal(c.Body(), &well_drill_int_detail); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_drill_int_detail values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39)")
	if err != nil {
		return err
	}
	well_drill_int_detail.Row_created_date = formatDate(well_drill_int_detail.Row_created_date)
	well_drill_int_detail.Row_changed_date = formatDate(well_drill_int_detail.Row_changed_date)
	well_drill_int_detail.Date_format_desc = formatDateString(well_drill_int_detail.Date_format_desc)
	well_drill_int_detail.Effective_date = formatDateString(well_drill_int_detail.Effective_date)
	well_drill_int_detail.End_date = formatDateString(well_drill_int_detail.End_date)
	well_drill_int_detail.End_time = formatDateString(well_drill_int_detail.End_time)
	well_drill_int_detail.Expiry_date = formatDateString(well_drill_int_detail.Expiry_date)
	well_drill_int_detail.Start_date = formatDateString(well_drill_int_detail.Start_date)
	well_drill_int_detail.Start_time = formatDateString(well_drill_int_detail.Start_time)
	well_drill_int_detail.Row_effective_date = formatDateString(well_drill_int_detail.Row_effective_date)
	well_drill_int_detail.Row_expiry_date = formatDateString(well_drill_int_detail.Row_expiry_date)






	rows, err := stmt.Exec(well_drill_int_detail.Uwi, well_drill_int_detail.Bit_source, well_drill_int_detail.Bit_interval_obs_no, well_drill_int_detail.Bit_detail_type, well_drill_int_detail.Detail_obs_no, well_drill_int_detail.Active_ind, well_drill_int_detail.Average_value, well_drill_int_detail.Average_value_ouom, well_drill_int_detail.Average_value_uom, well_drill_int_detail.Base_depth, well_drill_int_detail.Base_depth_ouom, well_drill_int_detail.Bit_detail_code, well_drill_int_detail.Date_format_desc, well_drill_int_detail.Detail_desc, well_drill_int_detail.Effective_date, well_drill_int_detail.End_date, well_drill_int_detail.End_time, well_drill_int_detail.Expiry_date, well_drill_int_detail.Max_value, well_drill_int_detail.Max_value_ouom, well_drill_int_detail.Max_value_uom, well_drill_int_detail.Min_value, well_drill_int_detail.Min_value_ouom, well_drill_int_detail.Min_value_uom, well_drill_int_detail.Ppdm_guid, well_drill_int_detail.Remark, well_drill_int_detail.Source, well_drill_int_detail.Start_date, well_drill_int_detail.Start_time, well_drill_int_detail.Timezone, well_drill_int_detail.Top_depth, well_drill_int_detail.Top_depth_ouom, well_drill_int_detail.Row_changed_by, well_drill_int_detail.Row_changed_date, well_drill_int_detail.Row_created_by, well_drill_int_detail.Row_created_date, well_drill_int_detail.Row_effective_date, well_drill_int_detail.Row_expiry_date, well_drill_int_detail.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellDrillIntDetail(c *fiber.Ctx) error {
	var well_drill_int_detail dto.Well_drill_int_detail

	setDefaults(&well_drill_int_detail)

	if err := json.Unmarshal(c.Body(), &well_drill_int_detail); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_drill_int_detail.Uwi = id

    if well_drill_int_detail.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_drill_int_detail set bit_source = :1, bit_interval_obs_no = :2, bit_detail_type = :3, detail_obs_no = :4, active_ind = :5, average_value = :6, average_value_ouom = :7, average_value_uom = :8, base_depth = :9, base_depth_ouom = :10, bit_detail_code = :11, date_format_desc = :12, detail_desc = :13, effective_date = :14, end_date = :15, end_time = :16, expiry_date = :17, max_value = :18, max_value_ouom = :19, max_value_uom = :20, min_value = :21, min_value_ouom = :22, min_value_uom = :23, ppdm_guid = :24, remark = :25, source = :26, start_date = :27, start_time = :28, timezone = :29, top_depth = :30, top_depth_ouom = :31, row_changed_by = :32, row_changed_date = :33, row_created_by = :34, row_effective_date = :35, row_expiry_date = :36, row_quality = :37 where uwi = :39")
	if err != nil {
		return err
	}

	well_drill_int_detail.Row_changed_date = formatDate(well_drill_int_detail.Row_changed_date)
	well_drill_int_detail.Date_format_desc = formatDateString(well_drill_int_detail.Date_format_desc)
	well_drill_int_detail.Effective_date = formatDateString(well_drill_int_detail.Effective_date)
	well_drill_int_detail.End_date = formatDateString(well_drill_int_detail.End_date)
	well_drill_int_detail.End_time = formatDateString(well_drill_int_detail.End_time)
	well_drill_int_detail.Expiry_date = formatDateString(well_drill_int_detail.Expiry_date)
	well_drill_int_detail.Start_date = formatDateString(well_drill_int_detail.Start_date)
	well_drill_int_detail.Start_time = formatDateString(well_drill_int_detail.Start_time)
	well_drill_int_detail.Row_effective_date = formatDateString(well_drill_int_detail.Row_effective_date)
	well_drill_int_detail.Row_expiry_date = formatDateString(well_drill_int_detail.Row_expiry_date)






	rows, err := stmt.Exec(well_drill_int_detail.Bit_source, well_drill_int_detail.Bit_interval_obs_no, well_drill_int_detail.Bit_detail_type, well_drill_int_detail.Detail_obs_no, well_drill_int_detail.Active_ind, well_drill_int_detail.Average_value, well_drill_int_detail.Average_value_ouom, well_drill_int_detail.Average_value_uom, well_drill_int_detail.Base_depth, well_drill_int_detail.Base_depth_ouom, well_drill_int_detail.Bit_detail_code, well_drill_int_detail.Date_format_desc, well_drill_int_detail.Detail_desc, well_drill_int_detail.Effective_date, well_drill_int_detail.End_date, well_drill_int_detail.End_time, well_drill_int_detail.Expiry_date, well_drill_int_detail.Max_value, well_drill_int_detail.Max_value_ouom, well_drill_int_detail.Max_value_uom, well_drill_int_detail.Min_value, well_drill_int_detail.Min_value_ouom, well_drill_int_detail.Min_value_uom, well_drill_int_detail.Ppdm_guid, well_drill_int_detail.Remark, well_drill_int_detail.Source, well_drill_int_detail.Start_date, well_drill_int_detail.Start_time, well_drill_int_detail.Timezone, well_drill_int_detail.Top_depth, well_drill_int_detail.Top_depth_ouom, well_drill_int_detail.Row_changed_by, well_drill_int_detail.Row_changed_date, well_drill_int_detail.Row_created_by, well_drill_int_detail.Row_effective_date, well_drill_int_detail.Row_expiry_date, well_drill_int_detail.Row_quality, well_drill_int_detail.Uwi)
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

func PatchWellDrillIntDetail(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_drill_int_detail set "
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
		} else if key == "date_format_desc" || key == "effective_date" || key == "end_date" || key == "end_time" || key == "expiry_date" || key == "start_date" || key == "start_time" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteWellDrillIntDetail(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_drill_int_detail dto.Well_drill_int_detail
	well_drill_int_detail.Uwi = id

	stmt, err := db.Prepare("delete from well_drill_int_detail where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_drill_int_detail.Uwi)
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


