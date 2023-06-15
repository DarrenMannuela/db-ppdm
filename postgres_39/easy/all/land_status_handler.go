package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetLandStatus(c *fiber.Ctx) error {
	rows, err := db.Query("select * from land_status")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Land_status

	for rows.Next() {
		var land_status dto.Land_status
		if err := rows.Scan(&land_status.Land_right_subtype, &land_status.Land_right_id, &land_status.Status_type, &land_status.Land_right_status, &land_status.Status_seq_no, &land_status.Active_ind, &land_status.Effective_date, &land_status.Effective_term, &land_status.Effective_term_ouom, &land_status.Expiry_date, &land_status.Ppdm_guid, &land_status.Reason, &land_status.Remark, &land_status.Section_of_act, &land_status.Section_of_act_date, &land_status.Source, &land_status.Undetermined_term_ind, &land_status.Row_changed_by, &land_status.Row_changed_date, &land_status.Row_created_by, &land_status.Row_created_date, &land_status.Row_effective_date, &land_status.Row_expiry_date, &land_status.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append land_status to result
		result = append(result, land_status)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetLandStatus(c *fiber.Ctx) error {
	var land_status dto.Land_status

	setDefaults(&land_status)

	if err := json.Unmarshal(c.Body(), &land_status); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into land_status values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24)")
	if err != nil {
		return err
	}
	land_status.Row_created_date = formatDate(land_status.Row_created_date)
	land_status.Row_changed_date = formatDate(land_status.Row_changed_date)
	land_status.Effective_date = formatDateString(land_status.Effective_date)
	land_status.Expiry_date = formatDateString(land_status.Expiry_date)
	land_status.Section_of_act_date = formatDateString(land_status.Section_of_act_date)
	land_status.Row_effective_date = formatDateString(land_status.Row_effective_date)
	land_status.Row_expiry_date = formatDateString(land_status.Row_expiry_date)






	rows, err := stmt.Exec(land_status.Land_right_subtype, land_status.Land_right_id, land_status.Status_type, land_status.Land_right_status, land_status.Status_seq_no, land_status.Active_ind, land_status.Effective_date, land_status.Effective_term, land_status.Effective_term_ouom, land_status.Expiry_date, land_status.Ppdm_guid, land_status.Reason, land_status.Remark, land_status.Section_of_act, land_status.Section_of_act_date, land_status.Source, land_status.Undetermined_term_ind, land_status.Row_changed_by, land_status.Row_changed_date, land_status.Row_created_by, land_status.Row_created_date, land_status.Row_effective_date, land_status.Row_expiry_date, land_status.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateLandStatus(c *fiber.Ctx) error {
	var land_status dto.Land_status

	setDefaults(&land_status)

	if err := json.Unmarshal(c.Body(), &land_status); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	land_status.Land_right_subtype = id

    if land_status.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update land_status set land_right_id = :1, status_type = :2, land_right_status = :3, status_seq_no = :4, active_ind = :5, effective_date = :6, effective_term = :7, effective_term_ouom = :8, expiry_date = :9, ppdm_guid = :10, reason = :11, remark = :12, section_of_act = :13, section_of_act_date = :14, source = :15, undetermined_term_ind = :16, row_changed_by = :17, row_changed_date = :18, row_created_by = :19, row_effective_date = :20, row_expiry_date = :21, row_quality = :22 where land_right_subtype = :24")
	if err != nil {
		return err
	}

	land_status.Row_changed_date = formatDate(land_status.Row_changed_date)
	land_status.Effective_date = formatDateString(land_status.Effective_date)
	land_status.Expiry_date = formatDateString(land_status.Expiry_date)
	land_status.Section_of_act_date = formatDateString(land_status.Section_of_act_date)
	land_status.Row_effective_date = formatDateString(land_status.Row_effective_date)
	land_status.Row_expiry_date = formatDateString(land_status.Row_expiry_date)






	rows, err := stmt.Exec(land_status.Land_right_id, land_status.Status_type, land_status.Land_right_status, land_status.Status_seq_no, land_status.Active_ind, land_status.Effective_date, land_status.Effective_term, land_status.Effective_term_ouom, land_status.Expiry_date, land_status.Ppdm_guid, land_status.Reason, land_status.Remark, land_status.Section_of_act, land_status.Section_of_act_date, land_status.Source, land_status.Undetermined_term_ind, land_status.Row_changed_by, land_status.Row_changed_date, land_status.Row_created_by, land_status.Row_effective_date, land_status.Row_expiry_date, land_status.Row_quality, land_status.Land_right_subtype)
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

func PatchLandStatus(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update land_status set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "section_of_act_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteLandStatus(c *fiber.Ctx) error {
	id := c.Params("id")
	var land_status dto.Land_status
	land_status.Land_right_subtype = id

	stmt, err := db.Prepare("delete from land_status where land_right_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(land_status.Land_right_subtype)
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


