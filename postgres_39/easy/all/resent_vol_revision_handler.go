package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetResentVolRevision(c *fiber.Ctx) error {
	rows, err := db.Query("select * from resent_vol_revision")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Resent_vol_revision

	for rows.Next() {
		var resent_vol_revision dto.Resent_vol_revision
		if err := rows.Scan(&resent_vol_revision.Resent_id, &resent_vol_revision.Reserve_class_id, &resent_vol_revision.Revision_id, &resent_vol_revision.Revision_obs_no, &resent_vol_revision.Active_ind, &resent_vol_revision.Analyst_ba_id, &resent_vol_revision.Approved_by_ba_id, &resent_vol_revision.Approved_date, &resent_vol_revision.Effective_date, &resent_vol_revision.Expiry_date, &resent_vol_revision.Gross_revision_ind, &resent_vol_revision.Interest_set_id, &resent_vol_revision.Interest_set_seq_no, &resent_vol_revision.Net_revision_ind, &resent_vol_revision.New_volume, &resent_vol_revision.New_volume_ouom, &resent_vol_revision.New_volume_uom, &resent_vol_revision.Partner_ba_id, &resent_vol_revision.Partner_obs_no, &resent_vol_revision.Ppdm_guid, &resent_vol_revision.Product_type, &resent_vol_revision.Project_id, &resent_vol_revision.Remark, &resent_vol_revision.Report_ind, &resent_vol_revision.Revision_category_id, &resent_vol_revision.Revision_date, &resent_vol_revision.Revision_date_desc, &resent_vol_revision.Revision_method, &resent_vol_revision.Revision_volume, &resent_vol_revision.Revision_volume_ouom, &resent_vol_revision.Revision_volume_uom, &resent_vol_revision.Source, &resent_vol_revision.Summary_id, &resent_vol_revision.Summary_obs_no, &resent_vol_revision.Row_changed_by, &resent_vol_revision.Row_changed_date, &resent_vol_revision.Row_created_by, &resent_vol_revision.Row_created_date, &resent_vol_revision.Row_effective_date, &resent_vol_revision.Row_expiry_date, &resent_vol_revision.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append resent_vol_revision to result
		result = append(result, resent_vol_revision)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetResentVolRevision(c *fiber.Ctx) error {
	var resent_vol_revision dto.Resent_vol_revision

	setDefaults(&resent_vol_revision)

	if err := json.Unmarshal(c.Body(), &resent_vol_revision); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into resent_vol_revision values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41)")
	if err != nil {
		return err
	}
	resent_vol_revision.Row_created_date = formatDate(resent_vol_revision.Row_created_date)
	resent_vol_revision.Row_changed_date = formatDate(resent_vol_revision.Row_changed_date)
	resent_vol_revision.Approved_date = formatDateString(resent_vol_revision.Approved_date)
	resent_vol_revision.Effective_date = formatDateString(resent_vol_revision.Effective_date)
	resent_vol_revision.Expiry_date = formatDateString(resent_vol_revision.Expiry_date)
	resent_vol_revision.Revision_date = formatDateString(resent_vol_revision.Revision_date)
	resent_vol_revision.Row_effective_date = formatDateString(resent_vol_revision.Row_effective_date)
	resent_vol_revision.Row_expiry_date = formatDateString(resent_vol_revision.Row_expiry_date)






	rows, err := stmt.Exec(resent_vol_revision.Resent_id, resent_vol_revision.Reserve_class_id, resent_vol_revision.Revision_id, resent_vol_revision.Revision_obs_no, resent_vol_revision.Active_ind, resent_vol_revision.Analyst_ba_id, resent_vol_revision.Approved_by_ba_id, resent_vol_revision.Approved_date, resent_vol_revision.Effective_date, resent_vol_revision.Expiry_date, resent_vol_revision.Gross_revision_ind, resent_vol_revision.Interest_set_id, resent_vol_revision.Interest_set_seq_no, resent_vol_revision.Net_revision_ind, resent_vol_revision.New_volume, resent_vol_revision.New_volume_ouom, resent_vol_revision.New_volume_uom, resent_vol_revision.Partner_ba_id, resent_vol_revision.Partner_obs_no, resent_vol_revision.Ppdm_guid, resent_vol_revision.Product_type, resent_vol_revision.Project_id, resent_vol_revision.Remark, resent_vol_revision.Report_ind, resent_vol_revision.Revision_category_id, resent_vol_revision.Revision_date, resent_vol_revision.Revision_date_desc, resent_vol_revision.Revision_method, resent_vol_revision.Revision_volume, resent_vol_revision.Revision_volume_ouom, resent_vol_revision.Revision_volume_uom, resent_vol_revision.Source, resent_vol_revision.Summary_id, resent_vol_revision.Summary_obs_no, resent_vol_revision.Row_changed_by, resent_vol_revision.Row_changed_date, resent_vol_revision.Row_created_by, resent_vol_revision.Row_created_date, resent_vol_revision.Row_effective_date, resent_vol_revision.Row_expiry_date, resent_vol_revision.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateResentVolRevision(c *fiber.Ctx) error {
	var resent_vol_revision dto.Resent_vol_revision

	setDefaults(&resent_vol_revision)

	if err := json.Unmarshal(c.Body(), &resent_vol_revision); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	resent_vol_revision.Resent_id = id

    if resent_vol_revision.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update resent_vol_revision set reserve_class_id = :1, revision_id = :2, revision_obs_no = :3, active_ind = :4, analyst_ba_id = :5, approved_by_ba_id = :6, approved_date = :7, effective_date = :8, expiry_date = :9, gross_revision_ind = :10, interest_set_id = :11, interest_set_seq_no = :12, net_revision_ind = :13, new_volume = :14, new_volume_ouom = :15, new_volume_uom = :16, partner_ba_id = :17, partner_obs_no = :18, ppdm_guid = :19, product_type = :20, project_id = :21, remark = :22, report_ind = :23, revision_category_id = :24, revision_date = :25, revision_date_desc = :26, revision_method = :27, revision_volume = :28, revision_volume_ouom = :29, revision_volume_uom = :30, source = :31, summary_id = :32, summary_obs_no = :33, row_changed_by = :34, row_changed_date = :35, row_created_by = :36, row_effective_date = :37, row_expiry_date = :38, row_quality = :39 where resent_id = :41")
	if err != nil {
		return err
	}

	resent_vol_revision.Row_changed_date = formatDate(resent_vol_revision.Row_changed_date)
	resent_vol_revision.Approved_date = formatDateString(resent_vol_revision.Approved_date)
	resent_vol_revision.Effective_date = formatDateString(resent_vol_revision.Effective_date)
	resent_vol_revision.Expiry_date = formatDateString(resent_vol_revision.Expiry_date)
	resent_vol_revision.Revision_date = formatDateString(resent_vol_revision.Revision_date)
	resent_vol_revision.Row_effective_date = formatDateString(resent_vol_revision.Row_effective_date)
	resent_vol_revision.Row_expiry_date = formatDateString(resent_vol_revision.Row_expiry_date)






	rows, err := stmt.Exec(resent_vol_revision.Reserve_class_id, resent_vol_revision.Revision_id, resent_vol_revision.Revision_obs_no, resent_vol_revision.Active_ind, resent_vol_revision.Analyst_ba_id, resent_vol_revision.Approved_by_ba_id, resent_vol_revision.Approved_date, resent_vol_revision.Effective_date, resent_vol_revision.Expiry_date, resent_vol_revision.Gross_revision_ind, resent_vol_revision.Interest_set_id, resent_vol_revision.Interest_set_seq_no, resent_vol_revision.Net_revision_ind, resent_vol_revision.New_volume, resent_vol_revision.New_volume_ouom, resent_vol_revision.New_volume_uom, resent_vol_revision.Partner_ba_id, resent_vol_revision.Partner_obs_no, resent_vol_revision.Ppdm_guid, resent_vol_revision.Product_type, resent_vol_revision.Project_id, resent_vol_revision.Remark, resent_vol_revision.Report_ind, resent_vol_revision.Revision_category_id, resent_vol_revision.Revision_date, resent_vol_revision.Revision_date_desc, resent_vol_revision.Revision_method, resent_vol_revision.Revision_volume, resent_vol_revision.Revision_volume_ouom, resent_vol_revision.Revision_volume_uom, resent_vol_revision.Source, resent_vol_revision.Summary_id, resent_vol_revision.Summary_obs_no, resent_vol_revision.Row_changed_by, resent_vol_revision.Row_changed_date, resent_vol_revision.Row_created_by, resent_vol_revision.Row_effective_date, resent_vol_revision.Row_expiry_date, resent_vol_revision.Row_quality, resent_vol_revision.Resent_id)
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

func PatchResentVolRevision(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update resent_vol_revision set "
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
		} else if key == "approved_date" || key == "effective_date" || key == "expiry_date" || key == "revision_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteResentVolRevision(c *fiber.Ctx) error {
	id := c.Params("id")
	var resent_vol_revision dto.Resent_vol_revision
	resent_vol_revision.Resent_id = id

	stmt, err := db.Prepare("delete from resent_vol_revision where resent_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(resent_vol_revision.Resent_id)
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


