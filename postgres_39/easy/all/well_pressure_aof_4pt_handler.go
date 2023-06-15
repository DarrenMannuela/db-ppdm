package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellPressureAof4Pt(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_pressure_aof_4pt")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_pressure_aof_4pt

	for rows.Next() {
		var well_pressure_aof_4pt dto.Well_pressure_aof_4pt
		if err := rows.Scan(&well_pressure_aof_4pt.Uwi, &well_pressure_aof_4pt.Source, &well_pressure_aof_4pt.Pressure_obs_no, &well_pressure_aof_4pt.Aof_obs_no, &well_pressure_aof_4pt.Point_obs_no, &well_pressure_aof_4pt.Active_ind, &well_pressure_aof_4pt.Bhfp, &well_pressure_aof_4pt.Bhfp_ouom, &well_pressure_aof_4pt.Choke_size_desc, &well_pressure_aof_4pt.Condensate_flow_rate, &well_pressure_aof_4pt.Condensate_flow_rate_ouom, &well_pressure_aof_4pt.Effective_date, &well_pressure_aof_4pt.Expiry_date, &well_pressure_aof_4pt.Final_shutin_pressure, &well_pressure_aof_4pt.Final_shutin_pressure_ouom, &well_pressure_aof_4pt.Flow_duration, &well_pressure_aof_4pt.Flow_duration_ouom, &well_pressure_aof_4pt.Flow_pressure, &well_pressure_aof_4pt.Flow_pressure_ouom, &well_pressure_aof_4pt.Flow_rate, &well_pressure_aof_4pt.Flow_rate_ouom, &well_pressure_aof_4pt.Ppdm_guid, &well_pressure_aof_4pt.Remark, &well_pressure_aof_4pt.Water_flow_rate, &well_pressure_aof_4pt.Water_flow_rate_ouom, &well_pressure_aof_4pt.Row_changed_by, &well_pressure_aof_4pt.Row_changed_date, &well_pressure_aof_4pt.Row_created_by, &well_pressure_aof_4pt.Row_created_date, &well_pressure_aof_4pt.Row_effective_date, &well_pressure_aof_4pt.Row_expiry_date, &well_pressure_aof_4pt.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_pressure_aof_4pt to result
		result = append(result, well_pressure_aof_4pt)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellPressureAof4Pt(c *fiber.Ctx) error {
	var well_pressure_aof_4pt dto.Well_pressure_aof_4pt

	setDefaults(&well_pressure_aof_4pt)

	if err := json.Unmarshal(c.Body(), &well_pressure_aof_4pt); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_pressure_aof_4pt values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32)")
	if err != nil {
		return err
	}
	well_pressure_aof_4pt.Row_created_date = formatDate(well_pressure_aof_4pt.Row_created_date)
	well_pressure_aof_4pt.Row_changed_date = formatDate(well_pressure_aof_4pt.Row_changed_date)
	well_pressure_aof_4pt.Effective_date = formatDateString(well_pressure_aof_4pt.Effective_date)
	well_pressure_aof_4pt.Expiry_date = formatDateString(well_pressure_aof_4pt.Expiry_date)
	well_pressure_aof_4pt.Row_effective_date = formatDateString(well_pressure_aof_4pt.Row_effective_date)
	well_pressure_aof_4pt.Row_expiry_date = formatDateString(well_pressure_aof_4pt.Row_expiry_date)






	rows, err := stmt.Exec(well_pressure_aof_4pt.Uwi, well_pressure_aof_4pt.Source, well_pressure_aof_4pt.Pressure_obs_no, well_pressure_aof_4pt.Aof_obs_no, well_pressure_aof_4pt.Point_obs_no, well_pressure_aof_4pt.Active_ind, well_pressure_aof_4pt.Bhfp, well_pressure_aof_4pt.Bhfp_ouom, well_pressure_aof_4pt.Choke_size_desc, well_pressure_aof_4pt.Condensate_flow_rate, well_pressure_aof_4pt.Condensate_flow_rate_ouom, well_pressure_aof_4pt.Effective_date, well_pressure_aof_4pt.Expiry_date, well_pressure_aof_4pt.Final_shutin_pressure, well_pressure_aof_4pt.Final_shutin_pressure_ouom, well_pressure_aof_4pt.Flow_duration, well_pressure_aof_4pt.Flow_duration_ouom, well_pressure_aof_4pt.Flow_pressure, well_pressure_aof_4pt.Flow_pressure_ouom, well_pressure_aof_4pt.Flow_rate, well_pressure_aof_4pt.Flow_rate_ouom, well_pressure_aof_4pt.Ppdm_guid, well_pressure_aof_4pt.Remark, well_pressure_aof_4pt.Water_flow_rate, well_pressure_aof_4pt.Water_flow_rate_ouom, well_pressure_aof_4pt.Row_changed_by, well_pressure_aof_4pt.Row_changed_date, well_pressure_aof_4pt.Row_created_by, well_pressure_aof_4pt.Row_created_date, well_pressure_aof_4pt.Row_effective_date, well_pressure_aof_4pt.Row_expiry_date, well_pressure_aof_4pt.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellPressureAof4Pt(c *fiber.Ctx) error {
	var well_pressure_aof_4pt dto.Well_pressure_aof_4pt

	setDefaults(&well_pressure_aof_4pt)

	if err := json.Unmarshal(c.Body(), &well_pressure_aof_4pt); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_pressure_aof_4pt.Uwi = id

    if well_pressure_aof_4pt.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_pressure_aof_4pt set source = :1, pressure_obs_no = :2, aof_obs_no = :3, point_obs_no = :4, active_ind = :5, bhfp = :6, bhfp_ouom = :7, choke_size_desc = :8, condensate_flow_rate = :9, condensate_flow_rate_ouom = :10, effective_date = :11, expiry_date = :12, final_shutin_pressure = :13, final_shutin_pressure_ouom = :14, flow_duration = :15, flow_duration_ouom = :16, flow_pressure = :17, flow_pressure_ouom = :18, flow_rate = :19, flow_rate_ouom = :20, ppdm_guid = :21, remark = :22, water_flow_rate = :23, water_flow_rate_ouom = :24, row_changed_by = :25, row_changed_date = :26, row_created_by = :27, row_effective_date = :28, row_expiry_date = :29, row_quality = :30 where uwi = :32")
	if err != nil {
		return err
	}

	well_pressure_aof_4pt.Row_changed_date = formatDate(well_pressure_aof_4pt.Row_changed_date)
	well_pressure_aof_4pt.Effective_date = formatDateString(well_pressure_aof_4pt.Effective_date)
	well_pressure_aof_4pt.Expiry_date = formatDateString(well_pressure_aof_4pt.Expiry_date)
	well_pressure_aof_4pt.Row_effective_date = formatDateString(well_pressure_aof_4pt.Row_effective_date)
	well_pressure_aof_4pt.Row_expiry_date = formatDateString(well_pressure_aof_4pt.Row_expiry_date)






	rows, err := stmt.Exec(well_pressure_aof_4pt.Source, well_pressure_aof_4pt.Pressure_obs_no, well_pressure_aof_4pt.Aof_obs_no, well_pressure_aof_4pt.Point_obs_no, well_pressure_aof_4pt.Active_ind, well_pressure_aof_4pt.Bhfp, well_pressure_aof_4pt.Bhfp_ouom, well_pressure_aof_4pt.Choke_size_desc, well_pressure_aof_4pt.Condensate_flow_rate, well_pressure_aof_4pt.Condensate_flow_rate_ouom, well_pressure_aof_4pt.Effective_date, well_pressure_aof_4pt.Expiry_date, well_pressure_aof_4pt.Final_shutin_pressure, well_pressure_aof_4pt.Final_shutin_pressure_ouom, well_pressure_aof_4pt.Flow_duration, well_pressure_aof_4pt.Flow_duration_ouom, well_pressure_aof_4pt.Flow_pressure, well_pressure_aof_4pt.Flow_pressure_ouom, well_pressure_aof_4pt.Flow_rate, well_pressure_aof_4pt.Flow_rate_ouom, well_pressure_aof_4pt.Ppdm_guid, well_pressure_aof_4pt.Remark, well_pressure_aof_4pt.Water_flow_rate, well_pressure_aof_4pt.Water_flow_rate_ouom, well_pressure_aof_4pt.Row_changed_by, well_pressure_aof_4pt.Row_changed_date, well_pressure_aof_4pt.Row_created_by, well_pressure_aof_4pt.Row_effective_date, well_pressure_aof_4pt.Row_expiry_date, well_pressure_aof_4pt.Row_quality, well_pressure_aof_4pt.Uwi)
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

func PatchWellPressureAof4Pt(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_pressure_aof_4pt set "
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

func DeleteWellPressureAof4Pt(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_pressure_aof_4pt dto.Well_pressure_aof_4pt
	well_pressure_aof_4pt.Uwi = id

	stmt, err := db.Prepare("delete from well_pressure_aof_4pt where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_pressure_aof_4pt.Uwi)
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


