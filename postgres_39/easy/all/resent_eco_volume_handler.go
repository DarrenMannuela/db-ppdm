package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetResentEcoVolume(c *fiber.Ctx) error {
	rows, err := db.Query("select * from resent_eco_volume")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Resent_eco_volume

	for rows.Next() {
		var resent_eco_volume dto.Resent_eco_volume
		if err := rows.Scan(&resent_eco_volume.Resent_id, &resent_eco_volume.Reserve_class_id, &resent_eco_volume.Economics_run_id, &resent_eco_volume.Product_type, &resent_eco_volume.Summary_obs_no, &resent_eco_volume.Active_ind, &resent_eco_volume.Effective_date, &resent_eco_volume.Expiry_date, &resent_eco_volume.Ppdm_guid, &resent_eco_volume.Remaining_balance, &resent_eco_volume.Remaining_balance_date, &resent_eco_volume.Remaining_balance_date_desc, &resent_eco_volume.Remaining_balance_ouom, &resent_eco_volume.Remaining_balance_uom, &resent_eco_volume.Remark, &resent_eco_volume.Source, &resent_eco_volume.Row_changed_by, &resent_eco_volume.Row_changed_date, &resent_eco_volume.Row_created_by, &resent_eco_volume.Row_created_date, &resent_eco_volume.Row_effective_date, &resent_eco_volume.Row_expiry_date, &resent_eco_volume.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append resent_eco_volume to result
		result = append(result, resent_eco_volume)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetResentEcoVolume(c *fiber.Ctx) error {
	var resent_eco_volume dto.Resent_eco_volume

	setDefaults(&resent_eco_volume)

	if err := json.Unmarshal(c.Body(), &resent_eco_volume); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into resent_eco_volume values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23)")
	if err != nil {
		return err
	}
	resent_eco_volume.Row_created_date = formatDate(resent_eco_volume.Row_created_date)
	resent_eco_volume.Row_changed_date = formatDate(resent_eco_volume.Row_changed_date)
	resent_eco_volume.Effective_date = formatDateString(resent_eco_volume.Effective_date)
	resent_eco_volume.Expiry_date = formatDateString(resent_eco_volume.Expiry_date)
	resent_eco_volume.Remaining_balance_date = formatDateString(resent_eco_volume.Remaining_balance_date)
	resent_eco_volume.Row_effective_date = formatDateString(resent_eco_volume.Row_effective_date)
	resent_eco_volume.Row_expiry_date = formatDateString(resent_eco_volume.Row_expiry_date)






	rows, err := stmt.Exec(resent_eco_volume.Resent_id, resent_eco_volume.Reserve_class_id, resent_eco_volume.Economics_run_id, resent_eco_volume.Product_type, resent_eco_volume.Summary_obs_no, resent_eco_volume.Active_ind, resent_eco_volume.Effective_date, resent_eco_volume.Expiry_date, resent_eco_volume.Ppdm_guid, resent_eco_volume.Remaining_balance, resent_eco_volume.Remaining_balance_date, resent_eco_volume.Remaining_balance_date_desc, resent_eco_volume.Remaining_balance_ouom, resent_eco_volume.Remaining_balance_uom, resent_eco_volume.Remark, resent_eco_volume.Source, resent_eco_volume.Row_changed_by, resent_eco_volume.Row_changed_date, resent_eco_volume.Row_created_by, resent_eco_volume.Row_created_date, resent_eco_volume.Row_effective_date, resent_eco_volume.Row_expiry_date, resent_eco_volume.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateResentEcoVolume(c *fiber.Ctx) error {
	var resent_eco_volume dto.Resent_eco_volume

	setDefaults(&resent_eco_volume)

	if err := json.Unmarshal(c.Body(), &resent_eco_volume); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	resent_eco_volume.Resent_id = id

    if resent_eco_volume.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update resent_eco_volume set reserve_class_id = :1, economics_run_id = :2, product_type = :3, summary_obs_no = :4, active_ind = :5, effective_date = :6, expiry_date = :7, ppdm_guid = :8, remaining_balance = :9, remaining_balance_date = :10, remaining_balance_date_desc = :11, remaining_balance_ouom = :12, remaining_balance_uom = :13, remark = :14, source = :15, row_changed_by = :16, row_changed_date = :17, row_created_by = :18, row_effective_date = :19, row_expiry_date = :20, row_quality = :21 where resent_id = :23")
	if err != nil {
		return err
	}

	resent_eco_volume.Row_changed_date = formatDate(resent_eco_volume.Row_changed_date)
	resent_eco_volume.Effective_date = formatDateString(resent_eco_volume.Effective_date)
	resent_eco_volume.Expiry_date = formatDateString(resent_eco_volume.Expiry_date)
	resent_eco_volume.Remaining_balance_date = formatDateString(resent_eco_volume.Remaining_balance_date)
	resent_eco_volume.Row_effective_date = formatDateString(resent_eco_volume.Row_effective_date)
	resent_eco_volume.Row_expiry_date = formatDateString(resent_eco_volume.Row_expiry_date)






	rows, err := stmt.Exec(resent_eco_volume.Reserve_class_id, resent_eco_volume.Economics_run_id, resent_eco_volume.Product_type, resent_eco_volume.Summary_obs_no, resent_eco_volume.Active_ind, resent_eco_volume.Effective_date, resent_eco_volume.Expiry_date, resent_eco_volume.Ppdm_guid, resent_eco_volume.Remaining_balance, resent_eco_volume.Remaining_balance_date, resent_eco_volume.Remaining_balance_date_desc, resent_eco_volume.Remaining_balance_ouom, resent_eco_volume.Remaining_balance_uom, resent_eco_volume.Remark, resent_eco_volume.Source, resent_eco_volume.Row_changed_by, resent_eco_volume.Row_changed_date, resent_eco_volume.Row_created_by, resent_eco_volume.Row_effective_date, resent_eco_volume.Row_expiry_date, resent_eco_volume.Row_quality, resent_eco_volume.Resent_id)
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

func PatchResentEcoVolume(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update resent_eco_volume set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "remaining_balance_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where resent_id = :id"

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

func DeleteResentEcoVolume(c *fiber.Ctx) error {
	id := c.Params("id")
	var resent_eco_volume dto.Resent_eco_volume
	resent_eco_volume.Resent_id = id

	stmt, err := db.Prepare("delete from resent_eco_volume where resent_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(resent_eco_volume.Resent_id)
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


