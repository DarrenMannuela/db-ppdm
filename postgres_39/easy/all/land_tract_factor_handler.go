package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetLandTractFactor(c *fiber.Ctx) error {
	rows, err := db.Query("select * from land_tract_factor")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Land_tract_factor

	for rows.Next() {
		var land_tract_factor dto.Land_tract_factor
		if err := rows.Scan(&land_tract_factor.Land_right_subtype, &land_tract_factor.Land_right_id, &land_tract_factor.Tract_factor_obs_no, &land_tract_factor.Active_ind, &land_tract_factor.Effective_date, &land_tract_factor.Expiry_date, &land_tract_factor.Ppdm_guid, &land_tract_factor.Remark, &land_tract_factor.Source, &land_tract_factor.Substance_id, &land_tract_factor.Tract_factor, &land_tract_factor.Row_changed_by, &land_tract_factor.Row_changed_date, &land_tract_factor.Row_created_by, &land_tract_factor.Row_created_date, &land_tract_factor.Row_effective_date, &land_tract_factor.Row_expiry_date, &land_tract_factor.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append land_tract_factor to result
		result = append(result, land_tract_factor)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetLandTractFactor(c *fiber.Ctx) error {
	var land_tract_factor dto.Land_tract_factor

	setDefaults(&land_tract_factor)

	if err := json.Unmarshal(c.Body(), &land_tract_factor); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into land_tract_factor values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18)")
	if err != nil {
		return err
	}
	land_tract_factor.Row_created_date = formatDate(land_tract_factor.Row_created_date)
	land_tract_factor.Row_changed_date = formatDate(land_tract_factor.Row_changed_date)
	land_tract_factor.Effective_date = formatDateString(land_tract_factor.Effective_date)
	land_tract_factor.Expiry_date = formatDateString(land_tract_factor.Expiry_date)
	land_tract_factor.Row_effective_date = formatDateString(land_tract_factor.Row_effective_date)
	land_tract_factor.Row_expiry_date = formatDateString(land_tract_factor.Row_expiry_date)






	rows, err := stmt.Exec(land_tract_factor.Land_right_subtype, land_tract_factor.Land_right_id, land_tract_factor.Tract_factor_obs_no, land_tract_factor.Active_ind, land_tract_factor.Effective_date, land_tract_factor.Expiry_date, land_tract_factor.Ppdm_guid, land_tract_factor.Remark, land_tract_factor.Source, land_tract_factor.Substance_id, land_tract_factor.Tract_factor, land_tract_factor.Row_changed_by, land_tract_factor.Row_changed_date, land_tract_factor.Row_created_by, land_tract_factor.Row_created_date, land_tract_factor.Row_effective_date, land_tract_factor.Row_expiry_date, land_tract_factor.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateLandTractFactor(c *fiber.Ctx) error {
	var land_tract_factor dto.Land_tract_factor

	setDefaults(&land_tract_factor)

	if err := json.Unmarshal(c.Body(), &land_tract_factor); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	land_tract_factor.Land_right_subtype = id

    if land_tract_factor.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update land_tract_factor set land_right_id = :1, tract_factor_obs_no = :2, active_ind = :3, effective_date = :4, expiry_date = :5, ppdm_guid = :6, remark = :7, source = :8, substance_id = :9, tract_factor = :10, row_changed_by = :11, row_changed_date = :12, row_created_by = :13, row_effective_date = :14, row_expiry_date = :15, row_quality = :16 where land_right_subtype = :18")
	if err != nil {
		return err
	}

	land_tract_factor.Row_changed_date = formatDate(land_tract_factor.Row_changed_date)
	land_tract_factor.Effective_date = formatDateString(land_tract_factor.Effective_date)
	land_tract_factor.Expiry_date = formatDateString(land_tract_factor.Expiry_date)
	land_tract_factor.Row_effective_date = formatDateString(land_tract_factor.Row_effective_date)
	land_tract_factor.Row_expiry_date = formatDateString(land_tract_factor.Row_expiry_date)






	rows, err := stmt.Exec(land_tract_factor.Land_right_id, land_tract_factor.Tract_factor_obs_no, land_tract_factor.Active_ind, land_tract_factor.Effective_date, land_tract_factor.Expiry_date, land_tract_factor.Ppdm_guid, land_tract_factor.Remark, land_tract_factor.Source, land_tract_factor.Substance_id, land_tract_factor.Tract_factor, land_tract_factor.Row_changed_by, land_tract_factor.Row_changed_date, land_tract_factor.Row_created_by, land_tract_factor.Row_effective_date, land_tract_factor.Row_expiry_date, land_tract_factor.Row_quality, land_tract_factor.Land_right_subtype)
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

func PatchLandTractFactor(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update land_tract_factor set "
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

func DeleteLandTractFactor(c *fiber.Ctx) error {
	id := c.Params("id")
	var land_tract_factor dto.Land_tract_factor
	land_tract_factor.Land_right_subtype = id

	stmt, err := db.Prepare("delete from land_tract_factor where land_right_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(land_tract_factor.Land_right_subtype)
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


