package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSeisVessel(c *fiber.Ctx) error {
	rows, err := db.Query("select * from seis_vessel")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Seis_vessel

	for rows.Next() {
		var seis_vessel dto.Seis_vessel
		if err := rows.Scan(&seis_vessel.Sf_subtype, &seis_vessel.Vessel_id, &seis_vessel.Vessel_config_obs_no, &seis_vessel.Acqtn_design_id, &seis_vessel.Active_ind, &seis_vessel.Effective_date, &seis_vessel.Expiry_date, &seis_vessel.Fath_azimuth, &seis_vessel.Fath_offset, &seis_vessel.Master_vessel_ind, &seis_vessel.Offset_ouom, &seis_vessel.Ppdm_guid, &seis_vessel.Reference_point, &seis_vessel.Remark, &seis_vessel.Shot_offset, &seis_vessel.Slave_vessel_ind, &seis_vessel.Source, &seis_vessel.Streamer_far_offset, &seis_vessel.Streamer_near_offset, &seis_vessel.Vessel_azimuth, &seis_vessel.Row_changed_by, &seis_vessel.Row_changed_date, &seis_vessel.Row_created_by, &seis_vessel.Row_created_date, &seis_vessel.Row_effective_date, &seis_vessel.Row_expiry_date, &seis_vessel.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append seis_vessel to result
		result = append(result, seis_vessel)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSeisVessel(c *fiber.Ctx) error {
	var seis_vessel dto.Seis_vessel

	setDefaults(&seis_vessel)

	if err := json.Unmarshal(c.Body(), &seis_vessel); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into seis_vessel values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27)")
	if err != nil {
		return err
	}
	seis_vessel.Row_created_date = formatDate(seis_vessel.Row_created_date)
	seis_vessel.Row_changed_date = formatDate(seis_vessel.Row_changed_date)
	seis_vessel.Effective_date = formatDateString(seis_vessel.Effective_date)
	seis_vessel.Expiry_date = formatDateString(seis_vessel.Expiry_date)
	seis_vessel.Row_effective_date = formatDateString(seis_vessel.Row_effective_date)
	seis_vessel.Row_expiry_date = formatDateString(seis_vessel.Row_expiry_date)






	rows, err := stmt.Exec(seis_vessel.Sf_subtype, seis_vessel.Vessel_id, seis_vessel.Vessel_config_obs_no, seis_vessel.Acqtn_design_id, seis_vessel.Active_ind, seis_vessel.Effective_date, seis_vessel.Expiry_date, seis_vessel.Fath_azimuth, seis_vessel.Fath_offset, seis_vessel.Master_vessel_ind, seis_vessel.Offset_ouom, seis_vessel.Ppdm_guid, seis_vessel.Reference_point, seis_vessel.Remark, seis_vessel.Shot_offset, seis_vessel.Slave_vessel_ind, seis_vessel.Source, seis_vessel.Streamer_far_offset, seis_vessel.Streamer_near_offset, seis_vessel.Vessel_azimuth, seis_vessel.Row_changed_by, seis_vessel.Row_changed_date, seis_vessel.Row_created_by, seis_vessel.Row_created_date, seis_vessel.Row_effective_date, seis_vessel.Row_expiry_date, seis_vessel.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSeisVessel(c *fiber.Ctx) error {
	var seis_vessel dto.Seis_vessel

	setDefaults(&seis_vessel)

	if err := json.Unmarshal(c.Body(), &seis_vessel); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	seis_vessel.Sf_subtype = id

    if seis_vessel.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update seis_vessel set vessel_id = :1, vessel_config_obs_no = :2, acqtn_design_id = :3, active_ind = :4, effective_date = :5, expiry_date = :6, fath_azimuth = :7, fath_offset = :8, master_vessel_ind = :9, offset_ouom = :10, ppdm_guid = :11, reference_point = :12, remark = :13, shot_offset = :14, slave_vessel_ind = :15, source = :16, streamer_far_offset = :17, streamer_near_offset = :18, vessel_azimuth = :19, row_changed_by = :20, row_changed_date = :21, row_created_by = :22, row_effective_date = :23, row_expiry_date = :24, row_quality = :25 where sf_subtype = :27")
	if err != nil {
		return err
	}

	seis_vessel.Row_changed_date = formatDate(seis_vessel.Row_changed_date)
	seis_vessel.Effective_date = formatDateString(seis_vessel.Effective_date)
	seis_vessel.Expiry_date = formatDateString(seis_vessel.Expiry_date)
	seis_vessel.Row_effective_date = formatDateString(seis_vessel.Row_effective_date)
	seis_vessel.Row_expiry_date = formatDateString(seis_vessel.Row_expiry_date)






	rows, err := stmt.Exec(seis_vessel.Vessel_id, seis_vessel.Vessel_config_obs_no, seis_vessel.Acqtn_design_id, seis_vessel.Active_ind, seis_vessel.Effective_date, seis_vessel.Expiry_date, seis_vessel.Fath_azimuth, seis_vessel.Fath_offset, seis_vessel.Master_vessel_ind, seis_vessel.Offset_ouom, seis_vessel.Ppdm_guid, seis_vessel.Reference_point, seis_vessel.Remark, seis_vessel.Shot_offset, seis_vessel.Slave_vessel_ind, seis_vessel.Source, seis_vessel.Streamer_far_offset, seis_vessel.Streamer_near_offset, seis_vessel.Vessel_azimuth, seis_vessel.Row_changed_by, seis_vessel.Row_changed_date, seis_vessel.Row_created_by, seis_vessel.Row_effective_date, seis_vessel.Row_expiry_date, seis_vessel.Row_quality, seis_vessel.Sf_subtype)
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

func PatchSeisVessel(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update seis_vessel set "
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
	query += " where sf_subtype = :id"

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

func DeleteSeisVessel(c *fiber.Ctx) error {
	id := c.Params("id")
	var seis_vessel dto.Seis_vessel
	seis_vessel.Sf_subtype = id

	stmt, err := db.Prepare("delete from seis_vessel where sf_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(seis_vessel.Sf_subtype)
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


