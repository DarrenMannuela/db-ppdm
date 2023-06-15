package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetApplication(c *fiber.Ctx) error {
	rows, err := db.Query("select * from application")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Application

	for rows.Next() {
		var application dto.Application
		if err := rows.Scan(&application.Application_id, &application.Active_ind, &application.Application_type, &application.Contract_id, &application.Current_status, &application.Decision, &application.Decision_date, &application.Effective_date, &application.Expiry_date, &application.Extension_id, &application.Fees_desc, &application.Fees_paid_ind, &application.Ppdm_guid, &application.Previous_application_id, &application.Rate_schedule_id, &application.Received_date, &application.Reference_num, &application.Remark, &application.Resubmission_ind, &application.Section_of_act, &application.Section_of_act_date, &application.Source, &application.Submission_complete_ind, &application.Submission_desc, &application.Submitted_by, &application.Submitted_date, &application.Submitted_to, &application.Row_changed_by, &application.Row_changed_date, &application.Row_created_by, &application.Row_created_date, &application.Row_effective_date, &application.Row_expiry_date, &application.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append application to result
		result = append(result, application)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetApplication(c *fiber.Ctx) error {
	var application dto.Application

	setDefaults(&application)

	if err := json.Unmarshal(c.Body(), &application); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into application values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34)")
	if err != nil {
		return err
	}
	application.Row_created_date = formatDate(application.Row_created_date)
	application.Row_changed_date = formatDate(application.Row_changed_date)
	application.Decision_date = formatDateString(application.Decision_date)
	application.Effective_date = formatDateString(application.Effective_date)
	application.Expiry_date = formatDateString(application.Expiry_date)
	application.Received_date = formatDateString(application.Received_date)
	application.Section_of_act_date = formatDateString(application.Section_of_act_date)
	application.Submitted_date = formatDateString(application.Submitted_date)
	application.Row_effective_date = formatDateString(application.Row_effective_date)
	application.Row_expiry_date = formatDateString(application.Row_expiry_date)






	rows, err := stmt.Exec(application.Application_id, application.Active_ind, application.Application_type, application.Contract_id, application.Current_status, application.Decision, application.Decision_date, application.Effective_date, application.Expiry_date, application.Extension_id, application.Fees_desc, application.Fees_paid_ind, application.Ppdm_guid, application.Previous_application_id, application.Rate_schedule_id, application.Received_date, application.Reference_num, application.Remark, application.Resubmission_ind, application.Section_of_act, application.Section_of_act_date, application.Source, application.Submission_complete_ind, application.Submission_desc, application.Submitted_by, application.Submitted_date, application.Submitted_to, application.Row_changed_by, application.Row_changed_date, application.Row_created_by, application.Row_created_date, application.Row_effective_date, application.Row_expiry_date, application.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateApplication(c *fiber.Ctx) error {
	var application dto.Application

	setDefaults(&application)

	if err := json.Unmarshal(c.Body(), &application); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	application.Application_id = id

    if application.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update application set active_ind = :1, application_type = :2, contract_id = :3, current_status = :4, decision = :5, decision_date = :6, effective_date = :7, expiry_date = :8, extension_id = :9, fees_desc = :10, fees_paid_ind = :11, ppdm_guid = :12, previous_application_id = :13, rate_schedule_id = :14, received_date = :15, reference_num = :16, remark = :17, resubmission_ind = :18, section_of_act = :19, section_of_act_date = :20, source = :21, submission_complete_ind = :22, submission_desc = :23, submitted_by = :24, submitted_date = :25, submitted_to = :26, row_changed_by = :27, row_changed_date = :28, row_created_by = :29, row_effective_date = :30, row_expiry_date = :31, row_quality = :32 where application_id = :34")
	if err != nil {
		return err
	}

	application.Row_changed_date = formatDate(application.Row_changed_date)
	application.Decision_date = formatDateString(application.Decision_date)
	application.Effective_date = formatDateString(application.Effective_date)
	application.Expiry_date = formatDateString(application.Expiry_date)
	application.Received_date = formatDateString(application.Received_date)
	application.Section_of_act_date = formatDateString(application.Section_of_act_date)
	application.Submitted_date = formatDateString(application.Submitted_date)
	application.Row_effective_date = formatDateString(application.Row_effective_date)
	application.Row_expiry_date = formatDateString(application.Row_expiry_date)






	rows, err := stmt.Exec(application.Active_ind, application.Application_type, application.Contract_id, application.Current_status, application.Decision, application.Decision_date, application.Effective_date, application.Expiry_date, application.Extension_id, application.Fees_desc, application.Fees_paid_ind, application.Ppdm_guid, application.Previous_application_id, application.Rate_schedule_id, application.Received_date, application.Reference_num, application.Remark, application.Resubmission_ind, application.Section_of_act, application.Section_of_act_date, application.Source, application.Submission_complete_ind, application.Submission_desc, application.Submitted_by, application.Submitted_date, application.Submitted_to, application.Row_changed_by, application.Row_changed_date, application.Row_created_by, application.Row_effective_date, application.Row_expiry_date, application.Row_quality, application.Application_id)
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

func PatchApplication(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update application set "
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
		} else if key == "decision_date" || key == "effective_date" || key == "expiry_date" || key == "received_date" || key == "section_of_act_date" || key == "submitted_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where application_id = :id"

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

func DeleteApplication(c *fiber.Ctx) error {
	id := c.Params("id")
	var application dto.Application
	application.Application_id = id

	stmt, err := db.Prepare("delete from application where application_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(application.Application_id)
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


