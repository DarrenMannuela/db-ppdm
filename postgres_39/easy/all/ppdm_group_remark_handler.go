package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPpdmGroupRemark(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ppdm_group_remark")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ppdm_group_remark

	for rows.Next() {
		var ppdm_group_remark dto.Ppdm_group_remark
		if err := rows.Scan(&ppdm_group_remark.Group_id, &ppdm_group_remark.Remark_obs_no, &ppdm_group_remark.Active_ind, &ppdm_group_remark.Effective_date, &ppdm_group_remark.Expiry_date, &ppdm_group_remark.Group_remark, &ppdm_group_remark.Made_by_ba_id, &ppdm_group_remark.Organization_id, &ppdm_group_remark.Organization_seq_no, &ppdm_group_remark.Ppdm_guid, &ppdm_group_remark.Remark, &ppdm_group_remark.Remark_type, &ppdm_group_remark.Source, &ppdm_group_remark.Row_changed_by, &ppdm_group_remark.Row_changed_date, &ppdm_group_remark.Row_created_by, &ppdm_group_remark.Row_created_date, &ppdm_group_remark.Row_effective_date, &ppdm_group_remark.Row_expiry_date, &ppdm_group_remark.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ppdm_group_remark to result
		result = append(result, ppdm_group_remark)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPpdmGroupRemark(c *fiber.Ctx) error {
	var ppdm_group_remark dto.Ppdm_group_remark

	setDefaults(&ppdm_group_remark)

	if err := json.Unmarshal(c.Body(), &ppdm_group_remark); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ppdm_group_remark values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20)")
	if err != nil {
		return err
	}
	ppdm_group_remark.Row_created_date = formatDate(ppdm_group_remark.Row_created_date)
	ppdm_group_remark.Row_changed_date = formatDate(ppdm_group_remark.Row_changed_date)
	ppdm_group_remark.Effective_date = formatDateString(ppdm_group_remark.Effective_date)
	ppdm_group_remark.Expiry_date = formatDateString(ppdm_group_remark.Expiry_date)
	ppdm_group_remark.Row_effective_date = formatDateString(ppdm_group_remark.Row_effective_date)
	ppdm_group_remark.Row_expiry_date = formatDateString(ppdm_group_remark.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_group_remark.Group_id, ppdm_group_remark.Remark_obs_no, ppdm_group_remark.Active_ind, ppdm_group_remark.Effective_date, ppdm_group_remark.Expiry_date, ppdm_group_remark.Group_remark, ppdm_group_remark.Made_by_ba_id, ppdm_group_remark.Organization_id, ppdm_group_remark.Organization_seq_no, ppdm_group_remark.Ppdm_guid, ppdm_group_remark.Remark, ppdm_group_remark.Remark_type, ppdm_group_remark.Source, ppdm_group_remark.Row_changed_by, ppdm_group_remark.Row_changed_date, ppdm_group_remark.Row_created_by, ppdm_group_remark.Row_created_date, ppdm_group_remark.Row_effective_date, ppdm_group_remark.Row_expiry_date, ppdm_group_remark.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePpdmGroupRemark(c *fiber.Ctx) error {
	var ppdm_group_remark dto.Ppdm_group_remark

	setDefaults(&ppdm_group_remark)

	if err := json.Unmarshal(c.Body(), &ppdm_group_remark); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ppdm_group_remark.Group_id = id

    if ppdm_group_remark.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ppdm_group_remark set remark_obs_no = :1, active_ind = :2, effective_date = :3, expiry_date = :4, group_remark = :5, made_by_ba_id = :6, organization_id = :7, organization_seq_no = :8, ppdm_guid = :9, remark = :10, remark_type = :11, source = :12, row_changed_by = :13, row_changed_date = :14, row_created_by = :15, row_effective_date = :16, row_expiry_date = :17, row_quality = :18 where group_id = :20")
	if err != nil {
		return err
	}

	ppdm_group_remark.Row_changed_date = formatDate(ppdm_group_remark.Row_changed_date)
	ppdm_group_remark.Effective_date = formatDateString(ppdm_group_remark.Effective_date)
	ppdm_group_remark.Expiry_date = formatDateString(ppdm_group_remark.Expiry_date)
	ppdm_group_remark.Row_effective_date = formatDateString(ppdm_group_remark.Row_effective_date)
	ppdm_group_remark.Row_expiry_date = formatDateString(ppdm_group_remark.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_group_remark.Remark_obs_no, ppdm_group_remark.Active_ind, ppdm_group_remark.Effective_date, ppdm_group_remark.Expiry_date, ppdm_group_remark.Group_remark, ppdm_group_remark.Made_by_ba_id, ppdm_group_remark.Organization_id, ppdm_group_remark.Organization_seq_no, ppdm_group_remark.Ppdm_guid, ppdm_group_remark.Remark, ppdm_group_remark.Remark_type, ppdm_group_remark.Source, ppdm_group_remark.Row_changed_by, ppdm_group_remark.Row_changed_date, ppdm_group_remark.Row_created_by, ppdm_group_remark.Row_effective_date, ppdm_group_remark.Row_expiry_date, ppdm_group_remark.Row_quality, ppdm_group_remark.Group_id)
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

func PatchPpdmGroupRemark(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ppdm_group_remark set "
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

func DeletePpdmGroupRemark(c *fiber.Ctx) error {
	id := c.Params("id")
	var ppdm_group_remark dto.Ppdm_group_remark
	ppdm_group_remark.Group_id = id

	stmt, err := db.Prepare("delete from ppdm_group_remark where group_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ppdm_group_remark.Group_id)
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


