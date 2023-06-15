package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetLandAgreePart(c *fiber.Ctx) error {
	rows, err := db.Query("select * from land_agree_part")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Land_agree_part

	for rows.Next() {
		var land_agree_part dto.Land_agree_part
		if err := rows.Scan(&land_agree_part.Land_right_subtype, &land_agree_part.Land_right_id, &land_agree_part.Active_ind, &land_agree_part.Dec_of_trust_ind, &land_agree_part.Dec_of_trust_remark, &land_agree_part.Effective_date, &land_agree_part.Expiry_date, &land_agree_part.Land_agree_type, &land_agree_part.Ppdm_guid, &land_agree_part.Remark, &land_agree_part.Rofr_ind, &land_agree_part.Source, &land_agree_part.Row_changed_by, &land_agree_part.Row_changed_date, &land_agree_part.Row_created_by, &land_agree_part.Row_created_date, &land_agree_part.Row_effective_date, &land_agree_part.Row_expiry_date, &land_agree_part.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append land_agree_part to result
		result = append(result, land_agree_part)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetLandAgreePart(c *fiber.Ctx) error {
	var land_agree_part dto.Land_agree_part

	setDefaults(&land_agree_part)

	if err := json.Unmarshal(c.Body(), &land_agree_part); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into land_agree_part values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19)")
	if err != nil {
		return err
	}
	land_agree_part.Row_created_date = formatDate(land_agree_part.Row_created_date)
	land_agree_part.Row_changed_date = formatDate(land_agree_part.Row_changed_date)
	land_agree_part.Effective_date = formatDateString(land_agree_part.Effective_date)
	land_agree_part.Expiry_date = formatDateString(land_agree_part.Expiry_date)
	land_agree_part.Row_effective_date = formatDateString(land_agree_part.Row_effective_date)
	land_agree_part.Row_expiry_date = formatDateString(land_agree_part.Row_expiry_date)






	rows, err := stmt.Exec(land_agree_part.Land_right_subtype, land_agree_part.Land_right_id, land_agree_part.Active_ind, land_agree_part.Dec_of_trust_ind, land_agree_part.Dec_of_trust_remark, land_agree_part.Effective_date, land_agree_part.Expiry_date, land_agree_part.Land_agree_type, land_agree_part.Ppdm_guid, land_agree_part.Remark, land_agree_part.Rofr_ind, land_agree_part.Source, land_agree_part.Row_changed_by, land_agree_part.Row_changed_date, land_agree_part.Row_created_by, land_agree_part.Row_created_date, land_agree_part.Row_effective_date, land_agree_part.Row_expiry_date, land_agree_part.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateLandAgreePart(c *fiber.Ctx) error {
	var land_agree_part dto.Land_agree_part

	setDefaults(&land_agree_part)

	if err := json.Unmarshal(c.Body(), &land_agree_part); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	land_agree_part.Land_right_subtype = id

    if land_agree_part.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update land_agree_part set land_right_id = :1, active_ind = :2, dec_of_trust_ind = :3, dec_of_trust_remark = :4, effective_date = :5, expiry_date = :6, land_agree_type = :7, ppdm_guid = :8, remark = :9, rofr_ind = :10, source = :11, row_changed_by = :12, row_changed_date = :13, row_created_by = :14, row_effective_date = :15, row_expiry_date = :16, row_quality = :17 where land_right_subtype = :19")
	if err != nil {
		return err
	}

	land_agree_part.Row_changed_date = formatDate(land_agree_part.Row_changed_date)
	land_agree_part.Effective_date = formatDateString(land_agree_part.Effective_date)
	land_agree_part.Expiry_date = formatDateString(land_agree_part.Expiry_date)
	land_agree_part.Row_effective_date = formatDateString(land_agree_part.Row_effective_date)
	land_agree_part.Row_expiry_date = formatDateString(land_agree_part.Row_expiry_date)






	rows, err := stmt.Exec(land_agree_part.Land_right_id, land_agree_part.Active_ind, land_agree_part.Dec_of_trust_ind, land_agree_part.Dec_of_trust_remark, land_agree_part.Effective_date, land_agree_part.Expiry_date, land_agree_part.Land_agree_type, land_agree_part.Ppdm_guid, land_agree_part.Remark, land_agree_part.Rofr_ind, land_agree_part.Source, land_agree_part.Row_changed_by, land_agree_part.Row_changed_date, land_agree_part.Row_created_by, land_agree_part.Row_effective_date, land_agree_part.Row_expiry_date, land_agree_part.Row_quality, land_agree_part.Land_right_subtype)
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

func PatchLandAgreePart(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update land_agree_part set "
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

func DeleteLandAgreePart(c *fiber.Ctx) error {
	id := c.Params("id")
	var land_agree_part dto.Land_agree_part
	land_agree_part.Land_right_subtype = id

	stmt, err := db.Prepare("delete from land_agree_part where land_right_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(land_agree_part.Land_right_subtype)
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


