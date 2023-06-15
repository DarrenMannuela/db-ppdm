package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSeisRecvrSetup(c *fiber.Ctx) error {
	rows, err := db.Query("select * from seis_recvr_setup")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Seis_recvr_setup

	for rows.Next() {
		var seis_recvr_setup dto.Seis_recvr_setup
		if err := rows.Scan(&seis_recvr_setup.Acqtn_design_id, &seis_recvr_setup.Rcvr_setup_id, &seis_recvr_setup.Active_ind, &seis_recvr_setup.Avg_feathering_angle, &seis_recvr_setup.Avg_streamer_depth, &seis_recvr_setup.Avg_streamer_depth_ouom, &seis_recvr_setup.Base_freq, &seis_recvr_setup.Depth_controller, &seis_recvr_setup.Effective_date, &seis_recvr_setup.Expiry_date, &seis_recvr_setup.Fixed_ind, &seis_recvr_setup.Group_spacing, &seis_recvr_setup.Group_spacing_ouom, &seis_recvr_setup.Inline_offset, &seis_recvr_setup.Inline_offset_direction, &seis_recvr_setup.Offline_offset, &seis_recvr_setup.Offline_offset_direction, &seis_recvr_setup.Offset_ouom, &seis_recvr_setup.Ppdm_guid, &seis_recvr_setup.Rcvr_array_type, &seis_recvr_setup.Rcvr_make, &seis_recvr_setup.Rcvr_phone_count, &seis_recvr_setup.Rcvr_spacing, &seis_recvr_setup.Rcvr_spacing_ouom, &seis_recvr_setup.Rcvr_type, &seis_recvr_setup.Remark, &seis_recvr_setup.Source, &seis_recvr_setup.Spread_description, &seis_recvr_setup.Spread_description_ouom, &seis_recvr_setup.Streamer_count, &seis_recvr_setup.Row_changed_by, &seis_recvr_setup.Row_changed_date, &seis_recvr_setup.Row_created_by, &seis_recvr_setup.Row_created_date, &seis_recvr_setup.Row_effective_date, &seis_recvr_setup.Row_expiry_date, &seis_recvr_setup.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append seis_recvr_setup to result
		result = append(result, seis_recvr_setup)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSeisRecvrSetup(c *fiber.Ctx) error {
	var seis_recvr_setup dto.Seis_recvr_setup

	setDefaults(&seis_recvr_setup)

	if err := json.Unmarshal(c.Body(), &seis_recvr_setup); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into seis_recvr_setup values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37)")
	if err != nil {
		return err
	}
	seis_recvr_setup.Row_created_date = formatDate(seis_recvr_setup.Row_created_date)
	seis_recvr_setup.Row_changed_date = formatDate(seis_recvr_setup.Row_changed_date)
	seis_recvr_setup.Effective_date = formatDateString(seis_recvr_setup.Effective_date)
	seis_recvr_setup.Expiry_date = formatDateString(seis_recvr_setup.Expiry_date)
	seis_recvr_setup.Row_effective_date = formatDateString(seis_recvr_setup.Row_effective_date)
	seis_recvr_setup.Row_expiry_date = formatDateString(seis_recvr_setup.Row_expiry_date)






	rows, err := stmt.Exec(seis_recvr_setup.Acqtn_design_id, seis_recvr_setup.Rcvr_setup_id, seis_recvr_setup.Active_ind, seis_recvr_setup.Avg_feathering_angle, seis_recvr_setup.Avg_streamer_depth, seis_recvr_setup.Avg_streamer_depth_ouom, seis_recvr_setup.Base_freq, seis_recvr_setup.Depth_controller, seis_recvr_setup.Effective_date, seis_recvr_setup.Expiry_date, seis_recvr_setup.Fixed_ind, seis_recvr_setup.Group_spacing, seis_recvr_setup.Group_spacing_ouom, seis_recvr_setup.Inline_offset, seis_recvr_setup.Inline_offset_direction, seis_recvr_setup.Offline_offset, seis_recvr_setup.Offline_offset_direction, seis_recvr_setup.Offset_ouom, seis_recvr_setup.Ppdm_guid, seis_recvr_setup.Rcvr_array_type, seis_recvr_setup.Rcvr_make, seis_recvr_setup.Rcvr_phone_count, seis_recvr_setup.Rcvr_spacing, seis_recvr_setup.Rcvr_spacing_ouom, seis_recvr_setup.Rcvr_type, seis_recvr_setup.Remark, seis_recvr_setup.Source, seis_recvr_setup.Spread_description, seis_recvr_setup.Spread_description_ouom, seis_recvr_setup.Streamer_count, seis_recvr_setup.Row_changed_by, seis_recvr_setup.Row_changed_date, seis_recvr_setup.Row_created_by, seis_recvr_setup.Row_created_date, seis_recvr_setup.Row_effective_date, seis_recvr_setup.Row_expiry_date, seis_recvr_setup.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSeisRecvrSetup(c *fiber.Ctx) error {
	var seis_recvr_setup dto.Seis_recvr_setup

	setDefaults(&seis_recvr_setup)

	if err := json.Unmarshal(c.Body(), &seis_recvr_setup); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	seis_recvr_setup.Acqtn_design_id = id

    if seis_recvr_setup.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update seis_recvr_setup set rcvr_setup_id = :1, active_ind = :2, avg_feathering_angle = :3, avg_streamer_depth = :4, avg_streamer_depth_ouom = :5, base_freq = :6, depth_controller = :7, effective_date = :8, expiry_date = :9, fixed_ind = :10, group_spacing = :11, group_spacing_ouom = :12, inline_offset = :13, inline_offset_direction = :14, offline_offset = :15, offline_offset_direction = :16, offset_ouom = :17, ppdm_guid = :18, rcvr_array_type = :19, rcvr_make = :20, rcvr_phone_count = :21, rcvr_spacing = :22, rcvr_spacing_ouom = :23, rcvr_type = :24, remark = :25, source = :26, spread_description = :27, spread_description_ouom = :28, streamer_count = :29, row_changed_by = :30, row_changed_date = :31, row_created_by = :32, row_effective_date = :33, row_expiry_date = :34, row_quality = :35 where acqtn_design_id = :37")
	if err != nil {
		return err
	}

	seis_recvr_setup.Row_changed_date = formatDate(seis_recvr_setup.Row_changed_date)
	seis_recvr_setup.Effective_date = formatDateString(seis_recvr_setup.Effective_date)
	seis_recvr_setup.Expiry_date = formatDateString(seis_recvr_setup.Expiry_date)
	seis_recvr_setup.Row_effective_date = formatDateString(seis_recvr_setup.Row_effective_date)
	seis_recvr_setup.Row_expiry_date = formatDateString(seis_recvr_setup.Row_expiry_date)






	rows, err := stmt.Exec(seis_recvr_setup.Rcvr_setup_id, seis_recvr_setup.Active_ind, seis_recvr_setup.Avg_feathering_angle, seis_recvr_setup.Avg_streamer_depth, seis_recvr_setup.Avg_streamer_depth_ouom, seis_recvr_setup.Base_freq, seis_recvr_setup.Depth_controller, seis_recvr_setup.Effective_date, seis_recvr_setup.Expiry_date, seis_recvr_setup.Fixed_ind, seis_recvr_setup.Group_spacing, seis_recvr_setup.Group_spacing_ouom, seis_recvr_setup.Inline_offset, seis_recvr_setup.Inline_offset_direction, seis_recvr_setup.Offline_offset, seis_recvr_setup.Offline_offset_direction, seis_recvr_setup.Offset_ouom, seis_recvr_setup.Ppdm_guid, seis_recvr_setup.Rcvr_array_type, seis_recvr_setup.Rcvr_make, seis_recvr_setup.Rcvr_phone_count, seis_recvr_setup.Rcvr_spacing, seis_recvr_setup.Rcvr_spacing_ouom, seis_recvr_setup.Rcvr_type, seis_recvr_setup.Remark, seis_recvr_setup.Source, seis_recvr_setup.Spread_description, seis_recvr_setup.Spread_description_ouom, seis_recvr_setup.Streamer_count, seis_recvr_setup.Row_changed_by, seis_recvr_setup.Row_changed_date, seis_recvr_setup.Row_created_by, seis_recvr_setup.Row_effective_date, seis_recvr_setup.Row_expiry_date, seis_recvr_setup.Row_quality, seis_recvr_setup.Acqtn_design_id)
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

func PatchSeisRecvrSetup(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update seis_recvr_setup set "
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
	query += " where acqtn_design_id = :id"

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

func DeleteSeisRecvrSetup(c *fiber.Ctx) error {
	id := c.Params("id")
	var seis_recvr_setup dto.Seis_recvr_setup
	seis_recvr_setup.Acqtn_design_id = id

	stmt, err := db.Prepare("delete from seis_recvr_setup where acqtn_design_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(seis_recvr_setup.Acqtn_design_id)
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


