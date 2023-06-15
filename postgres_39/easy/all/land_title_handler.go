package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetLandTitle(c *fiber.Ctx) error {
	rows, err := db.Query("select * from land_title")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Land_title

	for rows.Next() {
		var land_title dto.Land_title
		if err := rows.Scan(&land_title.Land_right_subtype, &land_title.Land_right_id, &land_title.Active_ind, &land_title.Certified_desc, &land_title.Effective_date, &land_title.Expiry_date, &land_title.Ppdm_guid, &land_title.Registration_date, &land_title.Remark, &land_title.Source, &land_title.Title_change_reason, &land_title.Title_holder, &land_title.Title_number, &land_title.Title_reference_num, &land_title.Title_type, &land_title.Row_changed_by, &land_title.Row_changed_date, &land_title.Row_created_by, &land_title.Row_created_date, &land_title.Row_effective_date, &land_title.Row_expiry_date, &land_title.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append land_title to result
		result = append(result, land_title)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetLandTitle(c *fiber.Ctx) error {
	var land_title dto.Land_title

	setDefaults(&land_title)

	if err := json.Unmarshal(c.Body(), &land_title); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into land_title values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22)")
	if err != nil {
		return err
	}
	land_title.Row_created_date = formatDate(land_title.Row_created_date)
	land_title.Row_changed_date = formatDate(land_title.Row_changed_date)
	land_title.Effective_date = formatDateString(land_title.Effective_date)
	land_title.Expiry_date = formatDateString(land_title.Expiry_date)
	land_title.Registration_date = formatDateString(land_title.Registration_date)
	land_title.Row_effective_date = formatDateString(land_title.Row_effective_date)
	land_title.Row_expiry_date = formatDateString(land_title.Row_expiry_date)






	rows, err := stmt.Exec(land_title.Land_right_subtype, land_title.Land_right_id, land_title.Active_ind, land_title.Certified_desc, land_title.Effective_date, land_title.Expiry_date, land_title.Ppdm_guid, land_title.Registration_date, land_title.Remark, land_title.Source, land_title.Title_change_reason, land_title.Title_holder, land_title.Title_number, land_title.Title_reference_num, land_title.Title_type, land_title.Row_changed_by, land_title.Row_changed_date, land_title.Row_created_by, land_title.Row_created_date, land_title.Row_effective_date, land_title.Row_expiry_date, land_title.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateLandTitle(c *fiber.Ctx) error {
	var land_title dto.Land_title

	setDefaults(&land_title)

	if err := json.Unmarshal(c.Body(), &land_title); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	land_title.Land_right_subtype = id

    if land_title.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update land_title set land_right_id = :1, active_ind = :2, certified_desc = :3, effective_date = :4, expiry_date = :5, ppdm_guid = :6, registration_date = :7, remark = :8, source = :9, title_change_reason = :10, title_holder = :11, title_number = :12, title_reference_num = :13, title_type = :14, row_changed_by = :15, row_changed_date = :16, row_created_by = :17, row_effective_date = :18, row_expiry_date = :19, row_quality = :20 where land_right_subtype = :22")
	if err != nil {
		return err
	}

	land_title.Row_changed_date = formatDate(land_title.Row_changed_date)
	land_title.Effective_date = formatDateString(land_title.Effective_date)
	land_title.Expiry_date = formatDateString(land_title.Expiry_date)
	land_title.Registration_date = formatDateString(land_title.Registration_date)
	land_title.Row_effective_date = formatDateString(land_title.Row_effective_date)
	land_title.Row_expiry_date = formatDateString(land_title.Row_expiry_date)






	rows, err := stmt.Exec(land_title.Land_right_id, land_title.Active_ind, land_title.Certified_desc, land_title.Effective_date, land_title.Expiry_date, land_title.Ppdm_guid, land_title.Registration_date, land_title.Remark, land_title.Source, land_title.Title_change_reason, land_title.Title_holder, land_title.Title_number, land_title.Title_reference_num, land_title.Title_type, land_title.Row_changed_by, land_title.Row_changed_date, land_title.Row_created_by, land_title.Row_effective_date, land_title.Row_expiry_date, land_title.Row_quality, land_title.Land_right_subtype)
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

func PatchLandTitle(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update land_title set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "registration_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteLandTitle(c *fiber.Ctx) error {
	id := c.Params("id")
	var land_title dto.Land_title
	land_title.Land_right_subtype = id

	stmt, err := db.Prepare("delete from land_title where land_right_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(land_title.Land_right_subtype)
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


