package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetInstrument(c *fiber.Ctx) error {
	rows, err := db.Query("select * from instrument")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Instrument

	for rows.Next() {
		var instrument dto.Instrument
		if err := rows.Scan(&instrument.Instrument_id, &instrument.Active_ind, &instrument.Add_for_service_ba_id, &instrument.Area_id, &instrument.Area_type, &instrument.Book_name, &instrument.Book_number, &instrument.Caveator_ba_id, &instrument.Completion_report_ind, &instrument.Description, &instrument.Discharge_date, &instrument.Discharge_ind, &instrument.Discharge_num, &instrument.Document_num, &instrument.Drilling_intent_ind, &instrument.Effective_date, &instrument.Expiry_date, &instrument.Instrument_type, &instrument.Jurisdiction, &instrument.Jurisdiction_add_obs_no, &instrument.Jurisdiction_add_source, &instrument.Line_number, &instrument.Page_number, &instrument.Plug_and_abandon_ind, &instrument.Ppdm_guid, &instrument.Production_lease_ind, &instrument.Received_date, &instrument.Recorded_date, &instrument.Registered_by_ba_id, &instrument.Registration_date, &instrument.Registration_num, &instrument.Remark, &instrument.Source, &instrument.Valid_lease_ind, &instrument.Row_changed_by, &instrument.Row_changed_date, &instrument.Row_created_by, &instrument.Row_created_date, &instrument.Row_effective_date, &instrument.Row_expiry_date, &instrument.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append instrument to result
		result = append(result, instrument)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetInstrument(c *fiber.Ctx) error {
	var instrument dto.Instrument

	setDefaults(&instrument)

	if err := json.Unmarshal(c.Body(), &instrument); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into instrument values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41)")
	if err != nil {
		return err
	}
	instrument.Row_created_date = formatDate(instrument.Row_created_date)
	instrument.Row_changed_date = formatDate(instrument.Row_changed_date)
	instrument.Discharge_date = formatDateString(instrument.Discharge_date)
	instrument.Effective_date = formatDateString(instrument.Effective_date)
	instrument.Expiry_date = formatDateString(instrument.Expiry_date)
	instrument.Received_date = formatDateString(instrument.Received_date)
	instrument.Recorded_date = formatDateString(instrument.Recorded_date)
	instrument.Registration_date = formatDateString(instrument.Registration_date)
	instrument.Row_effective_date = formatDateString(instrument.Row_effective_date)
	instrument.Row_expiry_date = formatDateString(instrument.Row_expiry_date)






	rows, err := stmt.Exec(instrument.Instrument_id, instrument.Active_ind, instrument.Add_for_service_ba_id, instrument.Area_id, instrument.Area_type, instrument.Book_name, instrument.Book_number, instrument.Caveator_ba_id, instrument.Completion_report_ind, instrument.Description, instrument.Discharge_date, instrument.Discharge_ind, instrument.Discharge_num, instrument.Document_num, instrument.Drilling_intent_ind, instrument.Effective_date, instrument.Expiry_date, instrument.Instrument_type, instrument.Jurisdiction, instrument.Jurisdiction_add_obs_no, instrument.Jurisdiction_add_source, instrument.Line_number, instrument.Page_number, instrument.Plug_and_abandon_ind, instrument.Ppdm_guid, instrument.Production_lease_ind, instrument.Received_date, instrument.Recorded_date, instrument.Registered_by_ba_id, instrument.Registration_date, instrument.Registration_num, instrument.Remark, instrument.Source, instrument.Valid_lease_ind, instrument.Row_changed_by, instrument.Row_changed_date, instrument.Row_created_by, instrument.Row_created_date, instrument.Row_effective_date, instrument.Row_expiry_date, instrument.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateInstrument(c *fiber.Ctx) error {
	var instrument dto.Instrument

	setDefaults(&instrument)

	if err := json.Unmarshal(c.Body(), &instrument); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	instrument.Instrument_id = id

    if instrument.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update instrument set active_ind = :1, add_for_service_ba_id = :2, area_id = :3, area_type = :4, book_name = :5, book_number = :6, caveator_ba_id = :7, completion_report_ind = :8, description = :9, discharge_date = :10, discharge_ind = :11, discharge_num = :12, document_num = :13, drilling_intent_ind = :14, effective_date = :15, expiry_date = :16, instrument_type = :17, jurisdiction = :18, jurisdiction_add_obs_no = :19, jurisdiction_add_source = :20, line_number = :21, page_number = :22, plug_and_abandon_ind = :23, ppdm_guid = :24, production_lease_ind = :25, received_date = :26, recorded_date = :27, registered_by_ba_id = :28, registration_date = :29, registration_num = :30, remark = :31, source = :32, valid_lease_ind = :33, row_changed_by = :34, row_changed_date = :35, row_created_by = :36, row_effective_date = :37, row_expiry_date = :38, row_quality = :39 where instrument_id = :41")
	if err != nil {
		return err
	}

	instrument.Row_changed_date = formatDate(instrument.Row_changed_date)
	instrument.Discharge_date = formatDateString(instrument.Discharge_date)
	instrument.Effective_date = formatDateString(instrument.Effective_date)
	instrument.Expiry_date = formatDateString(instrument.Expiry_date)
	instrument.Received_date = formatDateString(instrument.Received_date)
	instrument.Recorded_date = formatDateString(instrument.Recorded_date)
	instrument.Registration_date = formatDateString(instrument.Registration_date)
	instrument.Row_effective_date = formatDateString(instrument.Row_effective_date)
	instrument.Row_expiry_date = formatDateString(instrument.Row_expiry_date)






	rows, err := stmt.Exec(instrument.Active_ind, instrument.Add_for_service_ba_id, instrument.Area_id, instrument.Area_type, instrument.Book_name, instrument.Book_number, instrument.Caveator_ba_id, instrument.Completion_report_ind, instrument.Description, instrument.Discharge_date, instrument.Discharge_ind, instrument.Discharge_num, instrument.Document_num, instrument.Drilling_intent_ind, instrument.Effective_date, instrument.Expiry_date, instrument.Instrument_type, instrument.Jurisdiction, instrument.Jurisdiction_add_obs_no, instrument.Jurisdiction_add_source, instrument.Line_number, instrument.Page_number, instrument.Plug_and_abandon_ind, instrument.Ppdm_guid, instrument.Production_lease_ind, instrument.Received_date, instrument.Recorded_date, instrument.Registered_by_ba_id, instrument.Registration_date, instrument.Registration_num, instrument.Remark, instrument.Source, instrument.Valid_lease_ind, instrument.Row_changed_by, instrument.Row_changed_date, instrument.Row_created_by, instrument.Row_effective_date, instrument.Row_expiry_date, instrument.Row_quality, instrument.Instrument_id)
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

func PatchInstrument(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update instrument set "
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
		} else if key == "discharge_date" || key == "effective_date" || key == "expiry_date" || key == "received_date" || key == "recorded_date" || key == "registration_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where instrument_id = :id"

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

func DeleteInstrument(c *fiber.Ctx) error {
	id := c.Params("id")
	var instrument dto.Instrument
	instrument.Instrument_id = id

	stmt, err := db.Prepare("delete from instrument where instrument_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(instrument.Instrument_id)
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


