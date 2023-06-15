package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSpParcelRemark(c *fiber.Ctx) error {
	rows, err := db.Query("select * from sp_parcel_remark")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Sp_parcel_remark

	for rows.Next() {
		var sp_parcel_remark dto.Sp_parcel_remark
		if err := rows.Scan(&sp_parcel_remark.Spatial_description_id, &sp_parcel_remark.Spatial_obs_no, &sp_parcel_remark.Parcel_id, &sp_parcel_remark.Remark_seq_no, &sp_parcel_remark.Active_ind, &sp_parcel_remark.Effective_date, &sp_parcel_remark.Expiry_date, &sp_parcel_remark.Ppdm_guid, &sp_parcel_remark.Remark, &sp_parcel_remark.Remark_date, &sp_parcel_remark.Source, &sp_parcel_remark.Row_changed_by, &sp_parcel_remark.Row_changed_date, &sp_parcel_remark.Row_created_by, &sp_parcel_remark.Row_created_date, &sp_parcel_remark.Row_effective_date, &sp_parcel_remark.Row_expiry_date, &sp_parcel_remark.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append sp_parcel_remark to result
		result = append(result, sp_parcel_remark)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSpParcelRemark(c *fiber.Ctx) error {
	var sp_parcel_remark dto.Sp_parcel_remark

	setDefaults(&sp_parcel_remark)

	if err := json.Unmarshal(c.Body(), &sp_parcel_remark); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into sp_parcel_remark values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18)")
	if err != nil {
		return err
	}
	sp_parcel_remark.Row_created_date = formatDate(sp_parcel_remark.Row_created_date)
	sp_parcel_remark.Row_changed_date = formatDate(sp_parcel_remark.Row_changed_date)
	sp_parcel_remark.Effective_date = formatDateString(sp_parcel_remark.Effective_date)
	sp_parcel_remark.Expiry_date = formatDateString(sp_parcel_remark.Expiry_date)
	sp_parcel_remark.Remark_date = formatDateString(sp_parcel_remark.Remark_date)
	sp_parcel_remark.Row_effective_date = formatDateString(sp_parcel_remark.Row_effective_date)
	sp_parcel_remark.Row_expiry_date = formatDateString(sp_parcel_remark.Row_expiry_date)






	rows, err := stmt.Exec(sp_parcel_remark.Spatial_description_id, sp_parcel_remark.Spatial_obs_no, sp_parcel_remark.Parcel_id, sp_parcel_remark.Remark_seq_no, sp_parcel_remark.Active_ind, sp_parcel_remark.Effective_date, sp_parcel_remark.Expiry_date, sp_parcel_remark.Ppdm_guid, sp_parcel_remark.Remark, sp_parcel_remark.Remark_date, sp_parcel_remark.Source, sp_parcel_remark.Row_changed_by, sp_parcel_remark.Row_changed_date, sp_parcel_remark.Row_created_by, sp_parcel_remark.Row_created_date, sp_parcel_remark.Row_effective_date, sp_parcel_remark.Row_expiry_date, sp_parcel_remark.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSpParcelRemark(c *fiber.Ctx) error {
	var sp_parcel_remark dto.Sp_parcel_remark

	setDefaults(&sp_parcel_remark)

	if err := json.Unmarshal(c.Body(), &sp_parcel_remark); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	sp_parcel_remark.Spatial_description_id = id

    if sp_parcel_remark.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update sp_parcel_remark set spatial_obs_no = :1, parcel_id = :2, remark_seq_no = :3, active_ind = :4, effective_date = :5, expiry_date = :6, ppdm_guid = :7, remark = :8, remark_date = :9, source = :10, row_changed_by = :11, row_changed_date = :12, row_created_by = :13, row_effective_date = :14, row_expiry_date = :15, row_quality = :16 where spatial_description_id = :18")
	if err != nil {
		return err
	}

	sp_parcel_remark.Row_changed_date = formatDate(sp_parcel_remark.Row_changed_date)
	sp_parcel_remark.Effective_date = formatDateString(sp_parcel_remark.Effective_date)
	sp_parcel_remark.Expiry_date = formatDateString(sp_parcel_remark.Expiry_date)
	sp_parcel_remark.Remark_date = formatDateString(sp_parcel_remark.Remark_date)
	sp_parcel_remark.Row_effective_date = formatDateString(sp_parcel_remark.Row_effective_date)
	sp_parcel_remark.Row_expiry_date = formatDateString(sp_parcel_remark.Row_expiry_date)






	rows, err := stmt.Exec(sp_parcel_remark.Spatial_obs_no, sp_parcel_remark.Parcel_id, sp_parcel_remark.Remark_seq_no, sp_parcel_remark.Active_ind, sp_parcel_remark.Effective_date, sp_parcel_remark.Expiry_date, sp_parcel_remark.Ppdm_guid, sp_parcel_remark.Remark, sp_parcel_remark.Remark_date, sp_parcel_remark.Source, sp_parcel_remark.Row_changed_by, sp_parcel_remark.Row_changed_date, sp_parcel_remark.Row_created_by, sp_parcel_remark.Row_effective_date, sp_parcel_remark.Row_expiry_date, sp_parcel_remark.Row_quality, sp_parcel_remark.Spatial_description_id)
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

func PatchSpParcelRemark(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update sp_parcel_remark set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "remark_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where spatial_description_id = :id"

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

func DeleteSpParcelRemark(c *fiber.Ctx) error {
	id := c.Params("id")
	var sp_parcel_remark dto.Sp_parcel_remark
	sp_parcel_remark.Spatial_description_id = id

	stmt, err := db.Prepare("delete from sp_parcel_remark where spatial_description_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(sp_parcel_remark.Spatial_description_id)
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


