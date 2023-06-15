package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetLandSize(c *fiber.Ctx) error {
	rows, err := db.Query("select * from land_size")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Land_size

	for rows.Next() {
		var land_size dto.Land_size
		if err := rows.Scan(&land_size.Land_right_subtype, &land_size.Land_right_id, &land_size.Size_type, &land_size.Size_seq_no, &land_size.Active_ind, &land_size.Business_associate_id, &land_size.Effective_date, &land_size.Expiry_date, &land_size.Gross_size, &land_size.Interest_set_id, &land_size.Interest_set_seq_no, &land_size.Net_size, &land_size.Partner_obs_no, &land_size.Ppdm_guid, &land_size.Remark, &land_size.Size_ouom, &land_size.Source, &land_size.Row_changed_by, &land_size.Row_changed_date, &land_size.Row_created_by, &land_size.Row_created_date, &land_size.Row_effective_date, &land_size.Row_expiry_date, &land_size.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append land_size to result
		result = append(result, land_size)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetLandSize(c *fiber.Ctx) error {
	var land_size dto.Land_size

	setDefaults(&land_size)

	if err := json.Unmarshal(c.Body(), &land_size); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into land_size values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24)")
	if err != nil {
		return err
	}
	land_size.Row_created_date = formatDate(land_size.Row_created_date)
	land_size.Row_changed_date = formatDate(land_size.Row_changed_date)
	land_size.Effective_date = formatDateString(land_size.Effective_date)
	land_size.Expiry_date = formatDateString(land_size.Expiry_date)
	land_size.Row_effective_date = formatDateString(land_size.Row_effective_date)
	land_size.Row_expiry_date = formatDateString(land_size.Row_expiry_date)






	rows, err := stmt.Exec(land_size.Land_right_subtype, land_size.Land_right_id, land_size.Size_type, land_size.Size_seq_no, land_size.Active_ind, land_size.Business_associate_id, land_size.Effective_date, land_size.Expiry_date, land_size.Gross_size, land_size.Interest_set_id, land_size.Interest_set_seq_no, land_size.Net_size, land_size.Partner_obs_no, land_size.Ppdm_guid, land_size.Remark, land_size.Size_ouom, land_size.Source, land_size.Row_changed_by, land_size.Row_changed_date, land_size.Row_created_by, land_size.Row_created_date, land_size.Row_effective_date, land_size.Row_expiry_date, land_size.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateLandSize(c *fiber.Ctx) error {
	var land_size dto.Land_size

	setDefaults(&land_size)

	if err := json.Unmarshal(c.Body(), &land_size); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	land_size.Land_right_subtype = id

    if land_size.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update land_size set land_right_id = :1, size_type = :2, size_seq_no = :3, active_ind = :4, business_associate_id = :5, effective_date = :6, expiry_date = :7, gross_size = :8, interest_set_id = :9, interest_set_seq_no = :10, net_size = :11, partner_obs_no = :12, ppdm_guid = :13, remark = :14, size_ouom = :15, source = :16, row_changed_by = :17, row_changed_date = :18, row_created_by = :19, row_effective_date = :20, row_expiry_date = :21, row_quality = :22 where land_right_subtype = :24")
	if err != nil {
		return err
	}

	land_size.Row_changed_date = formatDate(land_size.Row_changed_date)
	land_size.Effective_date = formatDateString(land_size.Effective_date)
	land_size.Expiry_date = formatDateString(land_size.Expiry_date)
	land_size.Row_effective_date = formatDateString(land_size.Row_effective_date)
	land_size.Row_expiry_date = formatDateString(land_size.Row_expiry_date)






	rows, err := stmt.Exec(land_size.Land_right_id, land_size.Size_type, land_size.Size_seq_no, land_size.Active_ind, land_size.Business_associate_id, land_size.Effective_date, land_size.Expiry_date, land_size.Gross_size, land_size.Interest_set_id, land_size.Interest_set_seq_no, land_size.Net_size, land_size.Partner_obs_no, land_size.Ppdm_guid, land_size.Remark, land_size.Size_ouom, land_size.Source, land_size.Row_changed_by, land_size.Row_changed_date, land_size.Row_created_by, land_size.Row_effective_date, land_size.Row_expiry_date, land_size.Row_quality, land_size.Land_right_subtype)
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

func PatchLandSize(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update land_size set "
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

func DeleteLandSize(c *fiber.Ctx) error {
	id := c.Params("id")
	var land_size dto.Land_size
	land_size.Land_right_subtype = id

	stmt, err := db.Prepare("delete from land_size where land_right_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(land_size.Land_right_subtype)
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


