package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetLandOccupant(c *fiber.Ctx) error {
	rows, err := db.Query("select * from land_occupant")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Land_occupant

	for rows.Next() {
		var land_occupant dto.Land_occupant
		if err := rows.Scan(&land_occupant.Land_right_subtype, &land_occupant.Land_right_id, &land_occupant.Business_associate_id, &land_occupant.Occupancy_seq_no, &land_occupant.Active_ind, &land_occupant.Effective_date, &land_occupant.Expiry_date, &land_occupant.Ppdm_guid, &land_occupant.Remark, &land_occupant.Source, &land_occupant.Row_changed_by, &land_occupant.Row_changed_date, &land_occupant.Row_created_by, &land_occupant.Row_created_date, &land_occupant.Row_effective_date, &land_occupant.Row_expiry_date, &land_occupant.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append land_occupant to result
		result = append(result, land_occupant)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetLandOccupant(c *fiber.Ctx) error {
	var land_occupant dto.Land_occupant

	setDefaults(&land_occupant)

	if err := json.Unmarshal(c.Body(), &land_occupant); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into land_occupant values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17)")
	if err != nil {
		return err
	}
	land_occupant.Row_created_date = formatDate(land_occupant.Row_created_date)
	land_occupant.Row_changed_date = formatDate(land_occupant.Row_changed_date)
	land_occupant.Effective_date = formatDateString(land_occupant.Effective_date)
	land_occupant.Expiry_date = formatDateString(land_occupant.Expiry_date)
	land_occupant.Row_effective_date = formatDateString(land_occupant.Row_effective_date)
	land_occupant.Row_expiry_date = formatDateString(land_occupant.Row_expiry_date)






	rows, err := stmt.Exec(land_occupant.Land_right_subtype, land_occupant.Land_right_id, land_occupant.Business_associate_id, land_occupant.Occupancy_seq_no, land_occupant.Active_ind, land_occupant.Effective_date, land_occupant.Expiry_date, land_occupant.Ppdm_guid, land_occupant.Remark, land_occupant.Source, land_occupant.Row_changed_by, land_occupant.Row_changed_date, land_occupant.Row_created_by, land_occupant.Row_created_date, land_occupant.Row_effective_date, land_occupant.Row_expiry_date, land_occupant.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateLandOccupant(c *fiber.Ctx) error {
	var land_occupant dto.Land_occupant

	setDefaults(&land_occupant)

	if err := json.Unmarshal(c.Body(), &land_occupant); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	land_occupant.Land_right_subtype = id

    if land_occupant.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update land_occupant set land_right_id = :1, business_associate_id = :2, occupancy_seq_no = :3, active_ind = :4, effective_date = :5, expiry_date = :6, ppdm_guid = :7, remark = :8, source = :9, row_changed_by = :10, row_changed_date = :11, row_created_by = :12, row_effective_date = :13, row_expiry_date = :14, row_quality = :15 where land_right_subtype = :17")
	if err != nil {
		return err
	}

	land_occupant.Row_changed_date = formatDate(land_occupant.Row_changed_date)
	land_occupant.Effective_date = formatDateString(land_occupant.Effective_date)
	land_occupant.Expiry_date = formatDateString(land_occupant.Expiry_date)
	land_occupant.Row_effective_date = formatDateString(land_occupant.Row_effective_date)
	land_occupant.Row_expiry_date = formatDateString(land_occupant.Row_expiry_date)






	rows, err := stmt.Exec(land_occupant.Land_right_id, land_occupant.Business_associate_id, land_occupant.Occupancy_seq_no, land_occupant.Active_ind, land_occupant.Effective_date, land_occupant.Expiry_date, land_occupant.Ppdm_guid, land_occupant.Remark, land_occupant.Source, land_occupant.Row_changed_by, land_occupant.Row_changed_date, land_occupant.Row_created_by, land_occupant.Row_effective_date, land_occupant.Row_expiry_date, land_occupant.Row_quality, land_occupant.Land_right_subtype)
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

func PatchLandOccupant(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update land_occupant set "
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

func DeleteLandOccupant(c *fiber.Ctx) error {
	id := c.Params("id")
	var land_occupant dto.Land_occupant
	land_occupant.Land_right_subtype = id

	stmt, err := db.Prepare("delete from land_occupant where land_right_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(land_occupant.Land_right_subtype)
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


