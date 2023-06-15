package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSeisStreamer(c *fiber.Ctx) error {
	rows, err := db.Query("select * from seis_streamer")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Seis_streamer

	for rows.Next() {
		var seis_streamer dto.Seis_streamer
		if err := rows.Scan(&seis_streamer.Streamer_id, &seis_streamer.Acqtn_design_id, &seis_streamer.Active_ind, &seis_streamer.Cable_make, &seis_streamer.Description, &seis_streamer.Effective_date, &seis_streamer.Expiry_date, &seis_streamer.Ppdm_guid, &seis_streamer.Remark, &seis_streamer.Sf_subtype, &seis_streamer.Source, &seis_streamer.Streamer_length, &seis_streamer.Streamer_length_ouom, &seis_streamer.Streamer_position, &seis_streamer.Vessel_config_obs_no, &seis_streamer.Vessel_id, &seis_streamer.Row_changed_by, &seis_streamer.Row_changed_date, &seis_streamer.Row_created_by, &seis_streamer.Row_created_date, &seis_streamer.Row_effective_date, &seis_streamer.Row_expiry_date, &seis_streamer.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append seis_streamer to result
		result = append(result, seis_streamer)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSeisStreamer(c *fiber.Ctx) error {
	var seis_streamer dto.Seis_streamer

	setDefaults(&seis_streamer)

	if err := json.Unmarshal(c.Body(), &seis_streamer); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into seis_streamer values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23)")
	if err != nil {
		return err
	}
	seis_streamer.Row_created_date = formatDate(seis_streamer.Row_created_date)
	seis_streamer.Row_changed_date = formatDate(seis_streamer.Row_changed_date)
	seis_streamer.Effective_date = formatDateString(seis_streamer.Effective_date)
	seis_streamer.Expiry_date = formatDateString(seis_streamer.Expiry_date)
	seis_streamer.Row_effective_date = formatDateString(seis_streamer.Row_effective_date)
	seis_streamer.Row_expiry_date = formatDateString(seis_streamer.Row_expiry_date)






	rows, err := stmt.Exec(seis_streamer.Streamer_id, seis_streamer.Acqtn_design_id, seis_streamer.Active_ind, seis_streamer.Cable_make, seis_streamer.Description, seis_streamer.Effective_date, seis_streamer.Expiry_date, seis_streamer.Ppdm_guid, seis_streamer.Remark, seis_streamer.Sf_subtype, seis_streamer.Source, seis_streamer.Streamer_length, seis_streamer.Streamer_length_ouom, seis_streamer.Streamer_position, seis_streamer.Vessel_config_obs_no, seis_streamer.Vessel_id, seis_streamer.Row_changed_by, seis_streamer.Row_changed_date, seis_streamer.Row_created_by, seis_streamer.Row_created_date, seis_streamer.Row_effective_date, seis_streamer.Row_expiry_date, seis_streamer.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSeisStreamer(c *fiber.Ctx) error {
	var seis_streamer dto.Seis_streamer

	setDefaults(&seis_streamer)

	if err := json.Unmarshal(c.Body(), &seis_streamer); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	seis_streamer.Streamer_id = id

    if seis_streamer.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update seis_streamer set acqtn_design_id = :1, active_ind = :2, cable_make = :3, description = :4, effective_date = :5, expiry_date = :6, ppdm_guid = :7, remark = :8, sf_subtype = :9, source = :10, streamer_length = :11, streamer_length_ouom = :12, streamer_position = :13, vessel_config_obs_no = :14, vessel_id = :15, row_changed_by = :16, row_changed_date = :17, row_created_by = :18, row_effective_date = :19, row_expiry_date = :20, row_quality = :21 where streamer_id = :23")
	if err != nil {
		return err
	}

	seis_streamer.Row_changed_date = formatDate(seis_streamer.Row_changed_date)
	seis_streamer.Effective_date = formatDateString(seis_streamer.Effective_date)
	seis_streamer.Expiry_date = formatDateString(seis_streamer.Expiry_date)
	seis_streamer.Row_effective_date = formatDateString(seis_streamer.Row_effective_date)
	seis_streamer.Row_expiry_date = formatDateString(seis_streamer.Row_expiry_date)






	rows, err := stmt.Exec(seis_streamer.Acqtn_design_id, seis_streamer.Active_ind, seis_streamer.Cable_make, seis_streamer.Description, seis_streamer.Effective_date, seis_streamer.Expiry_date, seis_streamer.Ppdm_guid, seis_streamer.Remark, seis_streamer.Sf_subtype, seis_streamer.Source, seis_streamer.Streamer_length, seis_streamer.Streamer_length_ouom, seis_streamer.Streamer_position, seis_streamer.Vessel_config_obs_no, seis_streamer.Vessel_id, seis_streamer.Row_changed_by, seis_streamer.Row_changed_date, seis_streamer.Row_created_by, seis_streamer.Row_effective_date, seis_streamer.Row_expiry_date, seis_streamer.Row_quality, seis_streamer.Streamer_id)
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

func PatchSeisStreamer(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update seis_streamer set "
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
	query += " where streamer_id = :id"

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

func DeleteSeisStreamer(c *fiber.Ctx) error {
	id := c.Params("id")
	var seis_streamer dto.Seis_streamer
	seis_streamer.Streamer_id = id

	stmt, err := db.Prepare("delete from seis_streamer where streamer_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(seis_streamer.Streamer_id)
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


