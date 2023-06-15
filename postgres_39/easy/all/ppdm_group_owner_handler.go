package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPpdmGroupOwner(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ppdm_group_owner")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ppdm_group_owner

	for rows.Next() {
		var ppdm_group_owner dto.Ppdm_group_owner
		if err := rows.Scan(&ppdm_group_owner.Group_id, &ppdm_group_owner.Owner_obs_no, &ppdm_group_owner.Active_ind, &ppdm_group_owner.Default_unit_system_id, &ppdm_group_owner.Effective_date, &ppdm_group_owner.Expiry_date, &ppdm_group_owner.Organization_id, &ppdm_group_owner.Organization_seq_no, &ppdm_group_owner.Owner_ba_id, &ppdm_group_owner.Owner_role, &ppdm_group_owner.Ppdm_guid, &ppdm_group_owner.Remark, &ppdm_group_owner.Source, &ppdm_group_owner.Sw_application_id, &ppdm_group_owner.Row_changed_by, &ppdm_group_owner.Row_changed_date, &ppdm_group_owner.Row_created_by, &ppdm_group_owner.Row_created_date, &ppdm_group_owner.Row_effective_date, &ppdm_group_owner.Row_expiry_date, &ppdm_group_owner.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ppdm_group_owner to result
		result = append(result, ppdm_group_owner)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPpdmGroupOwner(c *fiber.Ctx) error {
	var ppdm_group_owner dto.Ppdm_group_owner

	setDefaults(&ppdm_group_owner)

	if err := json.Unmarshal(c.Body(), &ppdm_group_owner); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ppdm_group_owner values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21)")
	if err != nil {
		return err
	}
	ppdm_group_owner.Row_created_date = formatDate(ppdm_group_owner.Row_created_date)
	ppdm_group_owner.Row_changed_date = formatDate(ppdm_group_owner.Row_changed_date)
	ppdm_group_owner.Effective_date = formatDateString(ppdm_group_owner.Effective_date)
	ppdm_group_owner.Expiry_date = formatDateString(ppdm_group_owner.Expiry_date)
	ppdm_group_owner.Row_effective_date = formatDateString(ppdm_group_owner.Row_effective_date)
	ppdm_group_owner.Row_expiry_date = formatDateString(ppdm_group_owner.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_group_owner.Group_id, ppdm_group_owner.Owner_obs_no, ppdm_group_owner.Active_ind, ppdm_group_owner.Default_unit_system_id, ppdm_group_owner.Effective_date, ppdm_group_owner.Expiry_date, ppdm_group_owner.Organization_id, ppdm_group_owner.Organization_seq_no, ppdm_group_owner.Owner_ba_id, ppdm_group_owner.Owner_role, ppdm_group_owner.Ppdm_guid, ppdm_group_owner.Remark, ppdm_group_owner.Source, ppdm_group_owner.Sw_application_id, ppdm_group_owner.Row_changed_by, ppdm_group_owner.Row_changed_date, ppdm_group_owner.Row_created_by, ppdm_group_owner.Row_created_date, ppdm_group_owner.Row_effective_date, ppdm_group_owner.Row_expiry_date, ppdm_group_owner.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePpdmGroupOwner(c *fiber.Ctx) error {
	var ppdm_group_owner dto.Ppdm_group_owner

	setDefaults(&ppdm_group_owner)

	if err := json.Unmarshal(c.Body(), &ppdm_group_owner); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ppdm_group_owner.Group_id = id

    if ppdm_group_owner.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ppdm_group_owner set owner_obs_no = :1, active_ind = :2, default_unit_system_id = :3, effective_date = :4, expiry_date = :5, organization_id = :6, organization_seq_no = :7, owner_ba_id = :8, owner_role = :9, ppdm_guid = :10, remark = :11, source = :12, sw_application_id = :13, row_changed_by = :14, row_changed_date = :15, row_created_by = :16, row_effective_date = :17, row_expiry_date = :18, row_quality = :19 where group_id = :21")
	if err != nil {
		return err
	}

	ppdm_group_owner.Row_changed_date = formatDate(ppdm_group_owner.Row_changed_date)
	ppdm_group_owner.Effective_date = formatDateString(ppdm_group_owner.Effective_date)
	ppdm_group_owner.Expiry_date = formatDateString(ppdm_group_owner.Expiry_date)
	ppdm_group_owner.Row_effective_date = formatDateString(ppdm_group_owner.Row_effective_date)
	ppdm_group_owner.Row_expiry_date = formatDateString(ppdm_group_owner.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_group_owner.Owner_obs_no, ppdm_group_owner.Active_ind, ppdm_group_owner.Default_unit_system_id, ppdm_group_owner.Effective_date, ppdm_group_owner.Expiry_date, ppdm_group_owner.Organization_id, ppdm_group_owner.Organization_seq_no, ppdm_group_owner.Owner_ba_id, ppdm_group_owner.Owner_role, ppdm_group_owner.Ppdm_guid, ppdm_group_owner.Remark, ppdm_group_owner.Source, ppdm_group_owner.Sw_application_id, ppdm_group_owner.Row_changed_by, ppdm_group_owner.Row_changed_date, ppdm_group_owner.Row_created_by, ppdm_group_owner.Row_effective_date, ppdm_group_owner.Row_expiry_date, ppdm_group_owner.Row_quality, ppdm_group_owner.Group_id)
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

func PatchPpdmGroupOwner(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ppdm_group_owner set "
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
	query += " where group_id = :id"

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

func DeletePpdmGroupOwner(c *fiber.Ctx) error {
	id := c.Params("id")
	var ppdm_group_owner dto.Ppdm_group_owner
	ppdm_group_owner.Group_id = id

	stmt, err := db.Prepare("delete from ppdm_group_owner where group_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ppdm_group_owner.Group_id)
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


