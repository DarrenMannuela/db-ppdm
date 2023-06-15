package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPpdmGroupXref(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ppdm_group_xref")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ppdm_group_xref

	for rows.Next() {
		var ppdm_group_xref dto.Ppdm_group_xref
		if err := rows.Scan(&ppdm_group_xref.Parent_group_id, &ppdm_group_xref.Child_group_id, &ppdm_group_xref.Xref_obs_no, &ppdm_group_xref.Active_ind, &ppdm_group_xref.Effective_date, &ppdm_group_xref.Expiry_date, &ppdm_group_xref.Group_xref_type, &ppdm_group_xref.Owner_ba_id, &ppdm_group_xref.Ppdm_guid, &ppdm_group_xref.Remark, &ppdm_group_xref.Source, &ppdm_group_xref.Sw_application_id, &ppdm_group_xref.Row_changed_by, &ppdm_group_xref.Row_changed_date, &ppdm_group_xref.Row_created_by, &ppdm_group_xref.Row_created_date, &ppdm_group_xref.Row_effective_date, &ppdm_group_xref.Row_expiry_date, &ppdm_group_xref.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ppdm_group_xref to result
		result = append(result, ppdm_group_xref)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPpdmGroupXref(c *fiber.Ctx) error {
	var ppdm_group_xref dto.Ppdm_group_xref

	setDefaults(&ppdm_group_xref)

	if err := json.Unmarshal(c.Body(), &ppdm_group_xref); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ppdm_group_xref values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19)")
	if err != nil {
		return err
	}
	ppdm_group_xref.Row_created_date = formatDate(ppdm_group_xref.Row_created_date)
	ppdm_group_xref.Row_changed_date = formatDate(ppdm_group_xref.Row_changed_date)
	ppdm_group_xref.Effective_date = formatDateString(ppdm_group_xref.Effective_date)
	ppdm_group_xref.Expiry_date = formatDateString(ppdm_group_xref.Expiry_date)
	ppdm_group_xref.Row_effective_date = formatDateString(ppdm_group_xref.Row_effective_date)
	ppdm_group_xref.Row_expiry_date = formatDateString(ppdm_group_xref.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_group_xref.Parent_group_id, ppdm_group_xref.Child_group_id, ppdm_group_xref.Xref_obs_no, ppdm_group_xref.Active_ind, ppdm_group_xref.Effective_date, ppdm_group_xref.Expiry_date, ppdm_group_xref.Group_xref_type, ppdm_group_xref.Owner_ba_id, ppdm_group_xref.Ppdm_guid, ppdm_group_xref.Remark, ppdm_group_xref.Source, ppdm_group_xref.Sw_application_id, ppdm_group_xref.Row_changed_by, ppdm_group_xref.Row_changed_date, ppdm_group_xref.Row_created_by, ppdm_group_xref.Row_created_date, ppdm_group_xref.Row_effective_date, ppdm_group_xref.Row_expiry_date, ppdm_group_xref.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePpdmGroupXref(c *fiber.Ctx) error {
	var ppdm_group_xref dto.Ppdm_group_xref

	setDefaults(&ppdm_group_xref)

	if err := json.Unmarshal(c.Body(), &ppdm_group_xref); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ppdm_group_xref.Parent_group_id = id

    if ppdm_group_xref.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ppdm_group_xref set child_group_id = :1, xref_obs_no = :2, active_ind = :3, effective_date = :4, expiry_date = :5, group_xref_type = :6, owner_ba_id = :7, ppdm_guid = :8, remark = :9, source = :10, sw_application_id = :11, row_changed_by = :12, row_changed_date = :13, row_created_by = :14, row_effective_date = :15, row_expiry_date = :16, row_quality = :17 where parent_group_id = :19")
	if err != nil {
		return err
	}

	ppdm_group_xref.Row_changed_date = formatDate(ppdm_group_xref.Row_changed_date)
	ppdm_group_xref.Effective_date = formatDateString(ppdm_group_xref.Effective_date)
	ppdm_group_xref.Expiry_date = formatDateString(ppdm_group_xref.Expiry_date)
	ppdm_group_xref.Row_effective_date = formatDateString(ppdm_group_xref.Row_effective_date)
	ppdm_group_xref.Row_expiry_date = formatDateString(ppdm_group_xref.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_group_xref.Child_group_id, ppdm_group_xref.Xref_obs_no, ppdm_group_xref.Active_ind, ppdm_group_xref.Effective_date, ppdm_group_xref.Expiry_date, ppdm_group_xref.Group_xref_type, ppdm_group_xref.Owner_ba_id, ppdm_group_xref.Ppdm_guid, ppdm_group_xref.Remark, ppdm_group_xref.Source, ppdm_group_xref.Sw_application_id, ppdm_group_xref.Row_changed_by, ppdm_group_xref.Row_changed_date, ppdm_group_xref.Row_created_by, ppdm_group_xref.Row_effective_date, ppdm_group_xref.Row_expiry_date, ppdm_group_xref.Row_quality, ppdm_group_xref.Parent_group_id)
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

func PatchPpdmGroupXref(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ppdm_group_xref set "
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
	query += " where parent_group_id = :id"

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

func DeletePpdmGroupXref(c *fiber.Ctx) error {
	id := c.Params("id")
	var ppdm_group_xref dto.Ppdm_group_xref
	ppdm_group_xref.Parent_group_id = id

	stmt, err := db.Prepare("delete from ppdm_group_xref where parent_group_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ppdm_group_xref.Parent_group_id)
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


