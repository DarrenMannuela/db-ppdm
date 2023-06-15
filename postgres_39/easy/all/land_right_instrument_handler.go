package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetLandRightInstrument(c *fiber.Ctx) error {
	rows, err := db.Query("select * from land_right_instrument")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Land_right_instrument

	for rows.Next() {
		var land_right_instrument dto.Land_right_instrument
		if err := rows.Scan(&land_right_instrument.Land_right_subtype, &land_right_instrument.Land_right_id, &land_right_instrument.Instrument_id, &land_right_instrument.Lr_inst_seq_no, &land_right_instrument.Active_ind, &land_right_instrument.Discharge_date, &land_right_instrument.Discharge_ind, &land_right_instrument.Discharge_instrument_id, &land_right_instrument.Discharge_num, &land_right_instrument.Effective_date, &land_right_instrument.Expiry_date, &land_right_instrument.Ppdm_guid, &land_right_instrument.Remark, &land_right_instrument.Source, &land_right_instrument.Row_changed_by, &land_right_instrument.Row_changed_date, &land_right_instrument.Row_created_by, &land_right_instrument.Row_created_date, &land_right_instrument.Row_effective_date, &land_right_instrument.Row_expiry_date, &land_right_instrument.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append land_right_instrument to result
		result = append(result, land_right_instrument)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetLandRightInstrument(c *fiber.Ctx) error {
	var land_right_instrument dto.Land_right_instrument

	setDefaults(&land_right_instrument)

	if err := json.Unmarshal(c.Body(), &land_right_instrument); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into land_right_instrument values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21)")
	if err != nil {
		return err
	}
	land_right_instrument.Row_created_date = formatDate(land_right_instrument.Row_created_date)
	land_right_instrument.Row_changed_date = formatDate(land_right_instrument.Row_changed_date)
	land_right_instrument.Discharge_date = formatDateString(land_right_instrument.Discharge_date)
	land_right_instrument.Effective_date = formatDateString(land_right_instrument.Effective_date)
	land_right_instrument.Expiry_date = formatDateString(land_right_instrument.Expiry_date)
	land_right_instrument.Row_effective_date = formatDateString(land_right_instrument.Row_effective_date)
	land_right_instrument.Row_expiry_date = formatDateString(land_right_instrument.Row_expiry_date)






	rows, err := stmt.Exec(land_right_instrument.Land_right_subtype, land_right_instrument.Land_right_id, land_right_instrument.Instrument_id, land_right_instrument.Lr_inst_seq_no, land_right_instrument.Active_ind, land_right_instrument.Discharge_date, land_right_instrument.Discharge_ind, land_right_instrument.Discharge_instrument_id, land_right_instrument.Discharge_num, land_right_instrument.Effective_date, land_right_instrument.Expiry_date, land_right_instrument.Ppdm_guid, land_right_instrument.Remark, land_right_instrument.Source, land_right_instrument.Row_changed_by, land_right_instrument.Row_changed_date, land_right_instrument.Row_created_by, land_right_instrument.Row_created_date, land_right_instrument.Row_effective_date, land_right_instrument.Row_expiry_date, land_right_instrument.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateLandRightInstrument(c *fiber.Ctx) error {
	var land_right_instrument dto.Land_right_instrument

	setDefaults(&land_right_instrument)

	if err := json.Unmarshal(c.Body(), &land_right_instrument); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	land_right_instrument.Land_right_subtype = id

    if land_right_instrument.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update land_right_instrument set land_right_id = :1, instrument_id = :2, lr_inst_seq_no = :3, active_ind = :4, discharge_date = :5, discharge_ind = :6, discharge_instrument_id = :7, discharge_num = :8, effective_date = :9, expiry_date = :10, ppdm_guid = :11, remark = :12, source = :13, row_changed_by = :14, row_changed_date = :15, row_created_by = :16, row_effective_date = :17, row_expiry_date = :18, row_quality = :19 where land_right_subtype = :21")
	if err != nil {
		return err
	}

	land_right_instrument.Row_changed_date = formatDate(land_right_instrument.Row_changed_date)
	land_right_instrument.Discharge_date = formatDateString(land_right_instrument.Discharge_date)
	land_right_instrument.Effective_date = formatDateString(land_right_instrument.Effective_date)
	land_right_instrument.Expiry_date = formatDateString(land_right_instrument.Expiry_date)
	land_right_instrument.Row_effective_date = formatDateString(land_right_instrument.Row_effective_date)
	land_right_instrument.Row_expiry_date = formatDateString(land_right_instrument.Row_expiry_date)






	rows, err := stmt.Exec(land_right_instrument.Land_right_id, land_right_instrument.Instrument_id, land_right_instrument.Lr_inst_seq_no, land_right_instrument.Active_ind, land_right_instrument.Discharge_date, land_right_instrument.Discharge_ind, land_right_instrument.Discharge_instrument_id, land_right_instrument.Discharge_num, land_right_instrument.Effective_date, land_right_instrument.Expiry_date, land_right_instrument.Ppdm_guid, land_right_instrument.Remark, land_right_instrument.Source, land_right_instrument.Row_changed_by, land_right_instrument.Row_changed_date, land_right_instrument.Row_created_by, land_right_instrument.Row_effective_date, land_right_instrument.Row_expiry_date, land_right_instrument.Row_quality, land_right_instrument.Land_right_subtype)
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

func PatchLandRightInstrument(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update land_right_instrument set "
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
		} else if key == "discharge_date" || key == "effective_date" || key == "expiry_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where land_right_subtype = :id"

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

func DeleteLandRightInstrument(c *fiber.Ctx) error {
	id := c.Params("id")
	var land_right_instrument dto.Land_right_instrument
	land_right_instrument.Land_right_subtype = id

	stmt, err := db.Prepare("delete from land_right_instrument where land_right_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(land_right_instrument.Land_right_subtype)
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


