package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetLandArea(c *fiber.Ctx) error {
	rows, err := db.Query("select * from land_area")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Land_area

	for rows.Next() {
		var land_area dto.Land_area
		if err := rows.Scan(&land_area.Land_right_subtype, &land_area.Land_right_id, &land_area.Area_id, &land_area.Area_type, &land_area.Active_ind, &land_area.Effective_date, &land_area.Expiry_date, &land_area.Gross_size, &land_area.Gross_size_ouom, &land_area.Net_size, &land_area.Net_size_ouom, &land_area.Ppdm_guid, &land_area.Remark, &land_area.Source, &land_area.Row_changed_by, &land_area.Row_changed_date, &land_area.Row_created_by, &land_area.Row_created_date, &land_area.Row_effective_date, &land_area.Row_expiry_date, &land_area.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append land_area to result
		result = append(result, land_area)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetLandArea(c *fiber.Ctx) error {
	var land_area dto.Land_area

	setDefaults(&land_area)

	if err := json.Unmarshal(c.Body(), &land_area); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into land_area values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21)")
	if err != nil {
		return err
	}
	land_area.Row_created_date = formatDate(land_area.Row_created_date)
	land_area.Row_changed_date = formatDate(land_area.Row_changed_date)
	land_area.Effective_date = formatDateString(land_area.Effective_date)
	land_area.Expiry_date = formatDateString(land_area.Expiry_date)
	land_area.Row_effective_date = formatDateString(land_area.Row_effective_date)
	land_area.Row_expiry_date = formatDateString(land_area.Row_expiry_date)






	rows, err := stmt.Exec(land_area.Land_right_subtype, land_area.Land_right_id, land_area.Area_id, land_area.Area_type, land_area.Active_ind, land_area.Effective_date, land_area.Expiry_date, land_area.Gross_size, land_area.Gross_size_ouom, land_area.Net_size, land_area.Net_size_ouom, land_area.Ppdm_guid, land_area.Remark, land_area.Source, land_area.Row_changed_by, land_area.Row_changed_date, land_area.Row_created_by, land_area.Row_created_date, land_area.Row_effective_date, land_area.Row_expiry_date, land_area.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateLandArea(c *fiber.Ctx) error {
	var land_area dto.Land_area

	setDefaults(&land_area)

	if err := json.Unmarshal(c.Body(), &land_area); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	land_area.Land_right_subtype = id

    if land_area.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update land_area set land_right_id = :1, area_id = :2, area_type = :3, active_ind = :4, effective_date = :5, expiry_date = :6, gross_size = :7, gross_size_ouom = :8, net_size = :9, net_size_ouom = :10, ppdm_guid = :11, remark = :12, source = :13, row_changed_by = :14, row_changed_date = :15, row_created_by = :16, row_effective_date = :17, row_expiry_date = :18, row_quality = :19 where land_right_subtype = :21")
	if err != nil {
		return err
	}

	land_area.Row_changed_date = formatDate(land_area.Row_changed_date)
	land_area.Effective_date = formatDateString(land_area.Effective_date)
	land_area.Expiry_date = formatDateString(land_area.Expiry_date)
	land_area.Row_effective_date = formatDateString(land_area.Row_effective_date)
	land_area.Row_expiry_date = formatDateString(land_area.Row_expiry_date)






	rows, err := stmt.Exec(land_area.Land_right_id, land_area.Area_id, land_area.Area_type, land_area.Active_ind, land_area.Effective_date, land_area.Expiry_date, land_area.Gross_size, land_area.Gross_size_ouom, land_area.Net_size, land_area.Net_size_ouom, land_area.Ppdm_guid, land_area.Remark, land_area.Source, land_area.Row_changed_by, land_area.Row_changed_date, land_area.Row_created_by, land_area.Row_effective_date, land_area.Row_expiry_date, land_area.Row_quality, land_area.Land_right_subtype)
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

func PatchLandArea(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update land_area set "
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

func DeleteLandArea(c *fiber.Ctx) error {
	id := c.Params("id")
	var land_area dto.Land_area
	land_area.Land_right_subtype = id

	stmt, err := db.Prepare("delete from land_area where land_right_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(land_area.Land_right_subtype)
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


