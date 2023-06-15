package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetLandRemark(c *fiber.Ctx) error {
	rows, err := db.Query("select * from land_remark")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Land_remark

	for rows.Next() {
		var land_remark dto.Land_remark
		if err := rows.Scan(&land_remark.Land_right_subtype, &land_remark.Land_right_id, &land_remark.Remark_id, &land_remark.Remark_seq_no, &land_remark.Active_ind, &land_remark.Effective_date, &land_remark.Expiry_date, &land_remark.Ppdm_guid, &land_remark.Recommend_impl_ind, &land_remark.Remark, &land_remark.Remark_ba_id, &land_remark.Remark_date, &land_remark.Remark_type, &land_remark.Remark_user, &land_remark.Source, &land_remark.Row_changed_by, &land_remark.Row_changed_date, &land_remark.Row_created_by, &land_remark.Row_created_date, &land_remark.Row_effective_date, &land_remark.Row_expiry_date, &land_remark.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append land_remark to result
		result = append(result, land_remark)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetLandRemark(c *fiber.Ctx) error {
	var land_remark dto.Land_remark

	setDefaults(&land_remark)

	if err := json.Unmarshal(c.Body(), &land_remark); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into land_remark values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22)")
	if err != nil {
		return err
	}
	land_remark.Row_created_date = formatDate(land_remark.Row_created_date)
	land_remark.Row_changed_date = formatDate(land_remark.Row_changed_date)
	land_remark.Effective_date = formatDateString(land_remark.Effective_date)
	land_remark.Expiry_date = formatDateString(land_remark.Expiry_date)
	land_remark.Remark_date = formatDateString(land_remark.Remark_date)
	land_remark.Row_effective_date = formatDateString(land_remark.Row_effective_date)
	land_remark.Row_expiry_date = formatDateString(land_remark.Row_expiry_date)






	rows, err := stmt.Exec(land_remark.Land_right_subtype, land_remark.Land_right_id, land_remark.Remark_id, land_remark.Remark_seq_no, land_remark.Active_ind, land_remark.Effective_date, land_remark.Expiry_date, land_remark.Ppdm_guid, land_remark.Recommend_impl_ind, land_remark.Remark, land_remark.Remark_ba_id, land_remark.Remark_date, land_remark.Remark_type, land_remark.Remark_user, land_remark.Source, land_remark.Row_changed_by, land_remark.Row_changed_date, land_remark.Row_created_by, land_remark.Row_created_date, land_remark.Row_effective_date, land_remark.Row_expiry_date, land_remark.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateLandRemark(c *fiber.Ctx) error {
	var land_remark dto.Land_remark

	setDefaults(&land_remark)

	if err := json.Unmarshal(c.Body(), &land_remark); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	land_remark.Land_right_subtype = id

    if land_remark.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update land_remark set land_right_id = :1, remark_id = :2, remark_seq_no = :3, active_ind = :4, effective_date = :5, expiry_date = :6, ppdm_guid = :7, recommend_impl_ind = :8, remark = :9, remark_ba_id = :10, remark_date = :11, remark_type = :12, remark_user = :13, source = :14, row_changed_by = :15, row_changed_date = :16, row_created_by = :17, row_effective_date = :18, row_expiry_date = :19, row_quality = :20 where land_right_subtype = :22")
	if err != nil {
		return err
	}

	land_remark.Row_changed_date = formatDate(land_remark.Row_changed_date)
	land_remark.Effective_date = formatDateString(land_remark.Effective_date)
	land_remark.Expiry_date = formatDateString(land_remark.Expiry_date)
	land_remark.Remark_date = formatDateString(land_remark.Remark_date)
	land_remark.Row_effective_date = formatDateString(land_remark.Row_effective_date)
	land_remark.Row_expiry_date = formatDateString(land_remark.Row_expiry_date)






	rows, err := stmt.Exec(land_remark.Land_right_id, land_remark.Remark_id, land_remark.Remark_seq_no, land_remark.Active_ind, land_remark.Effective_date, land_remark.Expiry_date, land_remark.Ppdm_guid, land_remark.Recommend_impl_ind, land_remark.Remark, land_remark.Remark_ba_id, land_remark.Remark_date, land_remark.Remark_type, land_remark.Remark_user, land_remark.Source, land_remark.Row_changed_by, land_remark.Row_changed_date, land_remark.Row_created_by, land_remark.Row_effective_date, land_remark.Row_expiry_date, land_remark.Row_quality, land_remark.Land_right_subtype)
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

func PatchLandRemark(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update land_remark set "
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

func DeleteLandRemark(c *fiber.Ctx) error {
	id := c.Params("id")
	var land_remark dto.Land_remark
	land_remark.Land_right_subtype = id

	stmt, err := db.Prepare("delete from land_remark where land_right_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(land_remark.Land_right_subtype)
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


