package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetLithInterval(c *fiber.Ctx) error {
	rows, err := db.Query("select * from lith_interval")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Lith_interval

	for rows.Next() {
		var lith_interval dto.Lith_interval
		if err := rows.Scan(&lith_interval.Lithology_log_id, &lith_interval.Lith_log_source, &lith_interval.Depth_obs_no, &lith_interval.Active_ind, &lith_interval.Base_depth, &lith_interval.Base_depth_ouom, &lith_interval.Cycle_bed, &lith_interval.Effective_date, &lith_interval.Expiry_date, &lith_interval.No_data_desc, &lith_interval.Parent_depth_obs_no, &lith_interval.Ppdm_guid, &lith_interval.Preferred_ind, &lith_interval.Remark, &lith_interval.Sample_id, &lith_interval.Sample_quality, &lith_interval.Structure_type, &lith_interval.Thin_bed_ind, &lith_interval.Top_depth, &lith_interval.Top_depth_ouom, &lith_interval.Row_changed_by, &lith_interval.Row_changed_date, &lith_interval.Row_created_by, &lith_interval.Row_created_date, &lith_interval.Row_effective_date, &lith_interval.Row_expiry_date, &lith_interval.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append lith_interval to result
		result = append(result, lith_interval)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetLithInterval(c *fiber.Ctx) error {
	var lith_interval dto.Lith_interval

	setDefaults(&lith_interval)

	if err := json.Unmarshal(c.Body(), &lith_interval); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into lith_interval values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27)")
	if err != nil {
		return err
	}
	lith_interval.Row_created_date = formatDate(lith_interval.Row_created_date)
	lith_interval.Row_changed_date = formatDate(lith_interval.Row_changed_date)
	lith_interval.Effective_date = formatDateString(lith_interval.Effective_date)
	lith_interval.Expiry_date = formatDateString(lith_interval.Expiry_date)
	lith_interval.Row_effective_date = formatDateString(lith_interval.Row_effective_date)
	lith_interval.Row_expiry_date = formatDateString(lith_interval.Row_expiry_date)






	rows, err := stmt.Exec(lith_interval.Lithology_log_id, lith_interval.Lith_log_source, lith_interval.Depth_obs_no, lith_interval.Active_ind, lith_interval.Base_depth, lith_interval.Base_depth_ouom, lith_interval.Cycle_bed, lith_interval.Effective_date, lith_interval.Expiry_date, lith_interval.No_data_desc, lith_interval.Parent_depth_obs_no, lith_interval.Ppdm_guid, lith_interval.Preferred_ind, lith_interval.Remark, lith_interval.Sample_id, lith_interval.Sample_quality, lith_interval.Structure_type, lith_interval.Thin_bed_ind, lith_interval.Top_depth, lith_interval.Top_depth_ouom, lith_interval.Row_changed_by, lith_interval.Row_changed_date, lith_interval.Row_created_by, lith_interval.Row_created_date, lith_interval.Row_effective_date, lith_interval.Row_expiry_date, lith_interval.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateLithInterval(c *fiber.Ctx) error {
	var lith_interval dto.Lith_interval

	setDefaults(&lith_interval)

	if err := json.Unmarshal(c.Body(), &lith_interval); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	lith_interval.Lithology_log_id = id

    if lith_interval.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update lith_interval set lith_log_source = :1, depth_obs_no = :2, active_ind = :3, base_depth = :4, base_depth_ouom = :5, cycle_bed = :6, effective_date = :7, expiry_date = :8, no_data_desc = :9, parent_depth_obs_no = :10, ppdm_guid = :11, preferred_ind = :12, remark = :13, sample_id = :14, sample_quality = :15, structure_type = :16, thin_bed_ind = :17, top_depth = :18, top_depth_ouom = :19, row_changed_by = :20, row_changed_date = :21, row_created_by = :22, row_effective_date = :23, row_expiry_date = :24, row_quality = :25 where lithology_log_id = :27")
	if err != nil {
		return err
	}

	lith_interval.Row_changed_date = formatDate(lith_interval.Row_changed_date)
	lith_interval.Effective_date = formatDateString(lith_interval.Effective_date)
	lith_interval.Expiry_date = formatDateString(lith_interval.Expiry_date)
	lith_interval.Row_effective_date = formatDateString(lith_interval.Row_effective_date)
	lith_interval.Row_expiry_date = formatDateString(lith_interval.Row_expiry_date)






	rows, err := stmt.Exec(lith_interval.Lith_log_source, lith_interval.Depth_obs_no, lith_interval.Active_ind, lith_interval.Base_depth, lith_interval.Base_depth_ouom, lith_interval.Cycle_bed, lith_interval.Effective_date, lith_interval.Expiry_date, lith_interval.No_data_desc, lith_interval.Parent_depth_obs_no, lith_interval.Ppdm_guid, lith_interval.Preferred_ind, lith_interval.Remark, lith_interval.Sample_id, lith_interval.Sample_quality, lith_interval.Structure_type, lith_interval.Thin_bed_ind, lith_interval.Top_depth, lith_interval.Top_depth_ouom, lith_interval.Row_changed_by, lith_interval.Row_changed_date, lith_interval.Row_created_by, lith_interval.Row_effective_date, lith_interval.Row_expiry_date, lith_interval.Row_quality, lith_interval.Lithology_log_id)
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

func PatchLithInterval(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update lith_interval set "
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
	query += " where lithology_log_id = :id"

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

func DeleteLithInterval(c *fiber.Ctx) error {
	id := c.Params("id")
	var lith_interval dto.Lith_interval
	lith_interval.Lithology_log_id = id

	stmt, err := db.Prepare("delete from lith_interval where lithology_log_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(lith_interval.Lithology_log_id)
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


