package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetLandUnitTract(c *fiber.Ctx) error {
	rows, err := db.Query("select * from land_unit_tract")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Land_unit_tract

	for rows.Next() {
		var land_unit_tract dto.Land_unit_tract
		if err := rows.Scan(&land_unit_tract.Land_right_subtype, &land_unit_tract.Land_right_id, &land_unit_tract.Active_ind, &land_unit_tract.Effective_date, &land_unit_tract.Expiry_date, &land_unit_tract.Land_tract_type, &land_unit_tract.Land_unit_tract_name, &land_unit_tract.Land_unit_tract_number, &land_unit_tract.Percent_crown, &land_unit_tract.Percent_freehold, &land_unit_tract.Ppdm_guid, &land_unit_tract.Remark, &land_unit_tract.Source, &land_unit_tract.Row_changed_by, &land_unit_tract.Row_changed_date, &land_unit_tract.Row_created_by, &land_unit_tract.Row_created_date, &land_unit_tract.Row_effective_date, &land_unit_tract.Row_expiry_date, &land_unit_tract.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append land_unit_tract to result
		result = append(result, land_unit_tract)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetLandUnitTract(c *fiber.Ctx) error {
	var land_unit_tract dto.Land_unit_tract

	setDefaults(&land_unit_tract)

	if err := json.Unmarshal(c.Body(), &land_unit_tract); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into land_unit_tract values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20)")
	if err != nil {
		return err
	}
	land_unit_tract.Row_created_date = formatDate(land_unit_tract.Row_created_date)
	land_unit_tract.Row_changed_date = formatDate(land_unit_tract.Row_changed_date)
	land_unit_tract.Effective_date = formatDateString(land_unit_tract.Effective_date)
	land_unit_tract.Expiry_date = formatDateString(land_unit_tract.Expiry_date)
	land_unit_tract.Row_effective_date = formatDateString(land_unit_tract.Row_effective_date)
	land_unit_tract.Row_expiry_date = formatDateString(land_unit_tract.Row_expiry_date)






	rows, err := stmt.Exec(land_unit_tract.Land_right_subtype, land_unit_tract.Land_right_id, land_unit_tract.Active_ind, land_unit_tract.Effective_date, land_unit_tract.Expiry_date, land_unit_tract.Land_tract_type, land_unit_tract.Land_unit_tract_name, land_unit_tract.Land_unit_tract_number, land_unit_tract.Percent_crown, land_unit_tract.Percent_freehold, land_unit_tract.Ppdm_guid, land_unit_tract.Remark, land_unit_tract.Source, land_unit_tract.Row_changed_by, land_unit_tract.Row_changed_date, land_unit_tract.Row_created_by, land_unit_tract.Row_created_date, land_unit_tract.Row_effective_date, land_unit_tract.Row_expiry_date, land_unit_tract.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateLandUnitTract(c *fiber.Ctx) error {
	var land_unit_tract dto.Land_unit_tract

	setDefaults(&land_unit_tract)

	if err := json.Unmarshal(c.Body(), &land_unit_tract); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	land_unit_tract.Land_right_subtype = id

    if land_unit_tract.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update land_unit_tract set land_right_id = :1, active_ind = :2, effective_date = :3, expiry_date = :4, land_tract_type = :5, land_unit_tract_name = :6, land_unit_tract_number = :7, percent_crown = :8, percent_freehold = :9, ppdm_guid = :10, remark = :11, source = :12, row_changed_by = :13, row_changed_date = :14, row_created_by = :15, row_effective_date = :16, row_expiry_date = :17, row_quality = :18 where land_right_subtype = :20")
	if err != nil {
		return err
	}

	land_unit_tract.Row_changed_date = formatDate(land_unit_tract.Row_changed_date)
	land_unit_tract.Effective_date = formatDateString(land_unit_tract.Effective_date)
	land_unit_tract.Expiry_date = formatDateString(land_unit_tract.Expiry_date)
	land_unit_tract.Row_effective_date = formatDateString(land_unit_tract.Row_effective_date)
	land_unit_tract.Row_expiry_date = formatDateString(land_unit_tract.Row_expiry_date)






	rows, err := stmt.Exec(land_unit_tract.Land_right_id, land_unit_tract.Active_ind, land_unit_tract.Effective_date, land_unit_tract.Expiry_date, land_unit_tract.Land_tract_type, land_unit_tract.Land_unit_tract_name, land_unit_tract.Land_unit_tract_number, land_unit_tract.Percent_crown, land_unit_tract.Percent_freehold, land_unit_tract.Ppdm_guid, land_unit_tract.Remark, land_unit_tract.Source, land_unit_tract.Row_changed_by, land_unit_tract.Row_changed_date, land_unit_tract.Row_created_by, land_unit_tract.Row_effective_date, land_unit_tract.Row_expiry_date, land_unit_tract.Row_quality, land_unit_tract.Land_right_subtype)
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

func PatchLandUnitTract(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update land_unit_tract set "
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

func DeleteLandUnitTract(c *fiber.Ctx) error {
	id := c.Params("id")
	var land_unit_tract dto.Land_unit_tract
	land_unit_tract.Land_right_subtype = id

	stmt, err := db.Prepare("delete from land_unit_tract where land_right_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(land_unit_tract.Land_right_subtype)
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


