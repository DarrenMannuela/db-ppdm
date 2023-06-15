package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetLithRockStructure(c *fiber.Ctx) error {
	rows, err := db.Query("select * from lith_rock_structure")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Lith_rock_structure

	for rows.Next() {
		var lith_rock_structure dto.Lith_rock_structure
		if err := rows.Scan(&lith_rock_structure.Lithology_log_id, &lith_rock_structure.Lith_log_source, &lith_rock_structure.Depth_obs_no, &lith_rock_structure.Rock_type, &lith_rock_structure.Rock_type_obs_no, &lith_rock_structure.Structure_type, &lith_rock_structure.Active_ind, &lith_rock_structure.Boundary_type, &lith_rock_structure.Effective_date, &lith_rock_structure.Expiry_date, &lith_rock_structure.Ppdm_guid, &lith_rock_structure.Remark, &lith_rock_structure.Structure_rel_abund, &lith_rock_structure.Row_changed_by, &lith_rock_structure.Row_changed_date, &lith_rock_structure.Row_created_by, &lith_rock_structure.Row_created_date, &lith_rock_structure.Row_effective_date, &lith_rock_structure.Row_expiry_date, &lith_rock_structure.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append lith_rock_structure to result
		result = append(result, lith_rock_structure)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetLithRockStructure(c *fiber.Ctx) error {
	var lith_rock_structure dto.Lith_rock_structure

	setDefaults(&lith_rock_structure)

	if err := json.Unmarshal(c.Body(), &lith_rock_structure); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into lith_rock_structure values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20)")
	if err != nil {
		return err
	}
	lith_rock_structure.Row_created_date = formatDate(lith_rock_structure.Row_created_date)
	lith_rock_structure.Row_changed_date = formatDate(lith_rock_structure.Row_changed_date)
	lith_rock_structure.Effective_date = formatDateString(lith_rock_structure.Effective_date)
	lith_rock_structure.Expiry_date = formatDateString(lith_rock_structure.Expiry_date)
	lith_rock_structure.Row_effective_date = formatDateString(lith_rock_structure.Row_effective_date)
	lith_rock_structure.Row_expiry_date = formatDateString(lith_rock_structure.Row_expiry_date)






	rows, err := stmt.Exec(lith_rock_structure.Lithology_log_id, lith_rock_structure.Lith_log_source, lith_rock_structure.Depth_obs_no, lith_rock_structure.Rock_type, lith_rock_structure.Rock_type_obs_no, lith_rock_structure.Structure_type, lith_rock_structure.Active_ind, lith_rock_structure.Boundary_type, lith_rock_structure.Effective_date, lith_rock_structure.Expiry_date, lith_rock_structure.Ppdm_guid, lith_rock_structure.Remark, lith_rock_structure.Structure_rel_abund, lith_rock_structure.Row_changed_by, lith_rock_structure.Row_changed_date, lith_rock_structure.Row_created_by, lith_rock_structure.Row_created_date, lith_rock_structure.Row_effective_date, lith_rock_structure.Row_expiry_date, lith_rock_structure.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateLithRockStructure(c *fiber.Ctx) error {
	var lith_rock_structure dto.Lith_rock_structure

	setDefaults(&lith_rock_structure)

	if err := json.Unmarshal(c.Body(), &lith_rock_structure); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	lith_rock_structure.Lithology_log_id = id

    if lith_rock_structure.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update lith_rock_structure set lith_log_source = :1, depth_obs_no = :2, rock_type = :3, rock_type_obs_no = :4, structure_type = :5, active_ind = :6, boundary_type = :7, effective_date = :8, expiry_date = :9, ppdm_guid = :10, remark = :11, structure_rel_abund = :12, row_changed_by = :13, row_changed_date = :14, row_created_by = :15, row_effective_date = :16, row_expiry_date = :17, row_quality = :18 where lithology_log_id = :20")
	if err != nil {
		return err
	}

	lith_rock_structure.Row_changed_date = formatDate(lith_rock_structure.Row_changed_date)
	lith_rock_structure.Effective_date = formatDateString(lith_rock_structure.Effective_date)
	lith_rock_structure.Expiry_date = formatDateString(lith_rock_structure.Expiry_date)
	lith_rock_structure.Row_effective_date = formatDateString(lith_rock_structure.Row_effective_date)
	lith_rock_structure.Row_expiry_date = formatDateString(lith_rock_structure.Row_expiry_date)






	rows, err := stmt.Exec(lith_rock_structure.Lith_log_source, lith_rock_structure.Depth_obs_no, lith_rock_structure.Rock_type, lith_rock_structure.Rock_type_obs_no, lith_rock_structure.Structure_type, lith_rock_structure.Active_ind, lith_rock_structure.Boundary_type, lith_rock_structure.Effective_date, lith_rock_structure.Expiry_date, lith_rock_structure.Ppdm_guid, lith_rock_structure.Remark, lith_rock_structure.Structure_rel_abund, lith_rock_structure.Row_changed_by, lith_rock_structure.Row_changed_date, lith_rock_structure.Row_created_by, lith_rock_structure.Row_effective_date, lith_rock_structure.Row_expiry_date, lith_rock_structure.Row_quality, lith_rock_structure.Lithology_log_id)
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

func PatchLithRockStructure(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update lith_rock_structure set "
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

func DeleteLithRockStructure(c *fiber.Ctx) error {
	id := c.Params("id")
	var lith_rock_structure dto.Lith_rock_structure
	lith_rock_structure.Lithology_log_id = id

	stmt, err := db.Prepare("delete from lith_rock_structure where lithology_log_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(lith_rock_structure.Lithology_log_id)
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


