package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetLandXref(c *fiber.Ctx) error {
	rows, err := db.Query("select * from land_xref")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Land_xref

	for rows.Next() {
		var land_xref dto.Land_xref
		if err := rows.Scan(&land_xref.Parent_land_right_subtype, &land_xref.Child_land_right_subtype, &land_xref.Parent_land_right_id, &land_xref.Child_land_right_id, &land_xref.Lr_xref_seq_no, &land_xref.Active_ind, &land_xref.Effective_date, &land_xref.Expiry_date, &land_xref.Land_xref_type, &land_xref.Percent_allocation, &land_xref.Ppdm_guid, &land_xref.Remark, &land_xref.Source, &land_xref.Row_changed_by, &land_xref.Row_changed_date, &land_xref.Row_created_by, &land_xref.Row_created_date, &land_xref.Row_effective_date, &land_xref.Row_expiry_date, &land_xref.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append land_xref to result
		result = append(result, land_xref)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetLandXref(c *fiber.Ctx) error {
	var land_xref dto.Land_xref

	setDefaults(&land_xref)

	if err := json.Unmarshal(c.Body(), &land_xref); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into land_xref values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20)")
	if err != nil {
		return err
	}
	land_xref.Row_created_date = formatDate(land_xref.Row_created_date)
	land_xref.Row_changed_date = formatDate(land_xref.Row_changed_date)
	land_xref.Effective_date = formatDateString(land_xref.Effective_date)
	land_xref.Expiry_date = formatDateString(land_xref.Expiry_date)
	land_xref.Row_effective_date = formatDateString(land_xref.Row_effective_date)
	land_xref.Row_expiry_date = formatDateString(land_xref.Row_expiry_date)






	rows, err := stmt.Exec(land_xref.Parent_land_right_subtype, land_xref.Child_land_right_subtype, land_xref.Parent_land_right_id, land_xref.Child_land_right_id, land_xref.Lr_xref_seq_no, land_xref.Active_ind, land_xref.Effective_date, land_xref.Expiry_date, land_xref.Land_xref_type, land_xref.Percent_allocation, land_xref.Ppdm_guid, land_xref.Remark, land_xref.Source, land_xref.Row_changed_by, land_xref.Row_changed_date, land_xref.Row_created_by, land_xref.Row_created_date, land_xref.Row_effective_date, land_xref.Row_expiry_date, land_xref.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateLandXref(c *fiber.Ctx) error {
	var land_xref dto.Land_xref

	setDefaults(&land_xref)

	if err := json.Unmarshal(c.Body(), &land_xref); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	land_xref.Parent_land_right_subtype = id

    if land_xref.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update land_xref set child_land_right_subtype = :1, parent_land_right_id = :2, child_land_right_id = :3, lr_xref_seq_no = :4, active_ind = :5, effective_date = :6, expiry_date = :7, land_xref_type = :8, percent_allocation = :9, ppdm_guid = :10, remark = :11, source = :12, row_changed_by = :13, row_changed_date = :14, row_created_by = :15, row_effective_date = :16, row_expiry_date = :17, row_quality = :18 where parent_land_right_subtype = :20")
	if err != nil {
		return err
	}

	land_xref.Row_changed_date = formatDate(land_xref.Row_changed_date)
	land_xref.Effective_date = formatDateString(land_xref.Effective_date)
	land_xref.Expiry_date = formatDateString(land_xref.Expiry_date)
	land_xref.Row_effective_date = formatDateString(land_xref.Row_effective_date)
	land_xref.Row_expiry_date = formatDateString(land_xref.Row_expiry_date)






	rows, err := stmt.Exec(land_xref.Child_land_right_subtype, land_xref.Parent_land_right_id, land_xref.Child_land_right_id, land_xref.Lr_xref_seq_no, land_xref.Active_ind, land_xref.Effective_date, land_xref.Expiry_date, land_xref.Land_xref_type, land_xref.Percent_allocation, land_xref.Ppdm_guid, land_xref.Remark, land_xref.Source, land_xref.Row_changed_by, land_xref.Row_changed_date, land_xref.Row_created_by, land_xref.Row_effective_date, land_xref.Row_expiry_date, land_xref.Row_quality, land_xref.Parent_land_right_subtype)
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

func PatchLandXref(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update land_xref set "
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
	query += " where parent_land_right_subtype = :id"

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

func DeleteLandXref(c *fiber.Ctx) error {
	id := c.Params("id")
	var land_xref dto.Land_xref
	land_xref.Parent_land_right_subtype = id

	stmt, err := db.Prepare("delete from land_xref where parent_land_right_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(land_xref.Parent_land_right_subtype)
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


