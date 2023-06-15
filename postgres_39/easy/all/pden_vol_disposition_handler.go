package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPdenVolDisposition(c *fiber.Ctx) error {
	rows, err := db.Query("select * from pden_vol_disposition")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Pden_vol_disposition

	for rows.Next() {
		var pden_vol_disposition dto.Pden_vol_disposition
		if err := rows.Scan(&pden_vol_disposition.Pden_subtype, &pden_vol_disposition.Pden_id, &pden_vol_disposition.Pden_source, &pden_vol_disposition.Volume_method, &pden_vol_disposition.Activity_type, &pden_vol_disposition.Product_type, &pden_vol_disposition.Period_type, &pden_vol_disposition.Disposition_obs_no, &pden_vol_disposition.Amendment_seq_no, &pden_vol_disposition.Active_ind, &pden_vol_disposition.Amend_reason, &pden_vol_disposition.Business_associate_id, &pden_vol_disposition.Effective_date, &pden_vol_disposition.Expiry_date, &pden_vol_disposition.Posted_date, &pden_vol_disposition.Ppdm_guid, &pden_vol_disposition.Receiving_pden_id, &pden_vol_disposition.Receiving_pden_subtype, &pden_vol_disposition.Remark, &pden_vol_disposition.Report_ind, &pden_vol_disposition.Source, &pden_vol_disposition.Volume, &pden_vol_disposition.Volume_end_date, &pden_vol_disposition.Volume_end_date_desc, &pden_vol_disposition.Volume_ouom, &pden_vol_disposition.Volume_period, &pden_vol_disposition.Volume_period_ouom, &pden_vol_disposition.Volume_quality, &pden_vol_disposition.Volume_quality_ouom, &pden_vol_disposition.Volume_start_date, &pden_vol_disposition.Volume_uom, &pden_vol_disposition.Row_changed_by, &pden_vol_disposition.Row_changed_date, &pden_vol_disposition.Row_created_by, &pden_vol_disposition.Row_created_date, &pden_vol_disposition.Row_effective_date, &pden_vol_disposition.Row_expiry_date, &pden_vol_disposition.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append pden_vol_disposition to result
		result = append(result, pden_vol_disposition)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPdenVolDisposition(c *fiber.Ctx) error {
	var pden_vol_disposition dto.Pden_vol_disposition

	setDefaults(&pden_vol_disposition)

	if err := json.Unmarshal(c.Body(), &pden_vol_disposition); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into pden_vol_disposition values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38)")
	if err != nil {
		return err
	}
	pden_vol_disposition.Row_created_date = formatDate(pden_vol_disposition.Row_created_date)
	pden_vol_disposition.Row_changed_date = formatDate(pden_vol_disposition.Row_changed_date)
	pden_vol_disposition.Effective_date = formatDateString(pden_vol_disposition.Effective_date)
	pden_vol_disposition.Expiry_date = formatDateString(pden_vol_disposition.Expiry_date)
	pden_vol_disposition.Posted_date = formatDateString(pden_vol_disposition.Posted_date)
	pden_vol_disposition.Volume_end_date = formatDateString(pden_vol_disposition.Volume_end_date)
	pden_vol_disposition.Volume_start_date = formatDateString(pden_vol_disposition.Volume_start_date)
	pden_vol_disposition.Row_effective_date = formatDateString(pden_vol_disposition.Row_effective_date)
	pden_vol_disposition.Row_expiry_date = formatDateString(pden_vol_disposition.Row_expiry_date)






	rows, err := stmt.Exec(pden_vol_disposition.Pden_subtype, pden_vol_disposition.Pden_id, pden_vol_disposition.Pden_source, pden_vol_disposition.Volume_method, pden_vol_disposition.Activity_type, pden_vol_disposition.Product_type, pden_vol_disposition.Period_type, pden_vol_disposition.Disposition_obs_no, pden_vol_disposition.Amendment_seq_no, pden_vol_disposition.Active_ind, pden_vol_disposition.Amend_reason, pden_vol_disposition.Business_associate_id, pden_vol_disposition.Effective_date, pden_vol_disposition.Expiry_date, pden_vol_disposition.Posted_date, pden_vol_disposition.Ppdm_guid, pden_vol_disposition.Receiving_pden_id, pden_vol_disposition.Receiving_pden_subtype, pden_vol_disposition.Remark, pden_vol_disposition.Report_ind, pden_vol_disposition.Source, pden_vol_disposition.Volume, pden_vol_disposition.Volume_end_date, pden_vol_disposition.Volume_end_date_desc, pden_vol_disposition.Volume_ouom, pden_vol_disposition.Volume_period, pden_vol_disposition.Volume_period_ouom, pden_vol_disposition.Volume_quality, pden_vol_disposition.Volume_quality_ouom, pden_vol_disposition.Volume_start_date, pden_vol_disposition.Volume_uom, pden_vol_disposition.Row_changed_by, pden_vol_disposition.Row_changed_date, pden_vol_disposition.Row_created_by, pden_vol_disposition.Row_created_date, pden_vol_disposition.Row_effective_date, pden_vol_disposition.Row_expiry_date, pden_vol_disposition.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePdenVolDisposition(c *fiber.Ctx) error {
	var pden_vol_disposition dto.Pden_vol_disposition

	setDefaults(&pden_vol_disposition)

	if err := json.Unmarshal(c.Body(), &pden_vol_disposition); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	pden_vol_disposition.Pden_subtype = id

    if pden_vol_disposition.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update pden_vol_disposition set pden_id = :1, pden_source = :2, volume_method = :3, activity_type = :4, product_type = :5, period_type = :6, disposition_obs_no = :7, amendment_seq_no = :8, active_ind = :9, amend_reason = :10, business_associate_id = :11, effective_date = :12, expiry_date = :13, posted_date = :14, ppdm_guid = :15, receiving_pden_id = :16, receiving_pden_subtype = :17, remark = :18, report_ind = :19, source = :20, volume = :21, volume_end_date = :22, volume_end_date_desc = :23, volume_ouom = :24, volume_period = :25, volume_period_ouom = :26, volume_quality = :27, volume_quality_ouom = :28, volume_start_date = :29, volume_uom = :30, row_changed_by = :31, row_changed_date = :32, row_created_by = :33, row_effective_date = :34, row_expiry_date = :35, row_quality = :36 where pden_subtype = :38")
	if err != nil {
		return err
	}

	pden_vol_disposition.Row_changed_date = formatDate(pden_vol_disposition.Row_changed_date)
	pden_vol_disposition.Effective_date = formatDateString(pden_vol_disposition.Effective_date)
	pden_vol_disposition.Expiry_date = formatDateString(pden_vol_disposition.Expiry_date)
	pden_vol_disposition.Posted_date = formatDateString(pden_vol_disposition.Posted_date)
	pden_vol_disposition.Volume_end_date = formatDateString(pden_vol_disposition.Volume_end_date)
	pden_vol_disposition.Volume_start_date = formatDateString(pden_vol_disposition.Volume_start_date)
	pden_vol_disposition.Row_effective_date = formatDateString(pden_vol_disposition.Row_effective_date)
	pden_vol_disposition.Row_expiry_date = formatDateString(pden_vol_disposition.Row_expiry_date)






	rows, err := stmt.Exec(pden_vol_disposition.Pden_id, pden_vol_disposition.Pden_source, pden_vol_disposition.Volume_method, pden_vol_disposition.Activity_type, pden_vol_disposition.Product_type, pden_vol_disposition.Period_type, pden_vol_disposition.Disposition_obs_no, pden_vol_disposition.Amendment_seq_no, pden_vol_disposition.Active_ind, pden_vol_disposition.Amend_reason, pden_vol_disposition.Business_associate_id, pden_vol_disposition.Effective_date, pden_vol_disposition.Expiry_date, pden_vol_disposition.Posted_date, pden_vol_disposition.Ppdm_guid, pden_vol_disposition.Receiving_pden_id, pden_vol_disposition.Receiving_pden_subtype, pden_vol_disposition.Remark, pden_vol_disposition.Report_ind, pden_vol_disposition.Source, pden_vol_disposition.Volume, pden_vol_disposition.Volume_end_date, pden_vol_disposition.Volume_end_date_desc, pden_vol_disposition.Volume_ouom, pden_vol_disposition.Volume_period, pden_vol_disposition.Volume_period_ouom, pden_vol_disposition.Volume_quality, pden_vol_disposition.Volume_quality_ouom, pden_vol_disposition.Volume_start_date, pden_vol_disposition.Volume_uom, pden_vol_disposition.Row_changed_by, pden_vol_disposition.Row_changed_date, pden_vol_disposition.Row_created_by, pden_vol_disposition.Row_effective_date, pden_vol_disposition.Row_expiry_date, pden_vol_disposition.Row_quality, pden_vol_disposition.Pden_subtype)
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

func PatchPdenVolDisposition(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update pden_vol_disposition set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "posted_date" || key == "volume_end_date" || key == "volume_start_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where pden_subtype = :id"

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

func DeletePdenVolDisposition(c *fiber.Ctx) error {
	id := c.Params("id")
	var pden_vol_disposition dto.Pden_vol_disposition
	pden_vol_disposition.Pden_subtype = id

	stmt, err := db.Prepare("delete from pden_vol_disposition where pden_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(pden_vol_disposition.Pden_subtype)
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


