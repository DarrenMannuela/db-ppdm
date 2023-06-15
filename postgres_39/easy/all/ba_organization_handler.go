package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetBaOrganization(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ba_organization")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ba_organization

	for rows.Next() {
		var ba_organization dto.Ba_organization
		if err := rows.Scan(&ba_organization.Business_associate_id, &ba_organization.Organization_id, &ba_organization.Organization_seq_no, &ba_organization.Active_ind, &ba_organization.Address_obs_no, &ba_organization.Address_source, &ba_organization.Area_id, &ba_organization.Area_type, &ba_organization.Created_date, &ba_organization.Description, &ba_organization.Effective_date, &ba_organization.Expiry_date, &ba_organization.Organization_name, &ba_organization.Organization_type, &ba_organization.Ppdm_guid, &ba_organization.Remark, &ba_organization.Removed_date, &ba_organization.Source, &ba_organization.Row_changed_by, &ba_organization.Row_changed_date, &ba_organization.Row_created_by, &ba_organization.Row_created_date, &ba_organization.Row_effective_date, &ba_organization.Row_expiry_date, &ba_organization.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ba_organization to result
		result = append(result, ba_organization)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetBaOrganization(c *fiber.Ctx) error {
	var ba_organization dto.Ba_organization

	setDefaults(&ba_organization)

	if err := json.Unmarshal(c.Body(), &ba_organization); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ba_organization values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25)")
	if err != nil {
		return err
	}
	ba_organization.Row_created_date = formatDate(ba_organization.Row_created_date)
	ba_organization.Row_changed_date = formatDate(ba_organization.Row_changed_date)
	ba_organization.Created_date = formatDateString(ba_organization.Created_date)
	ba_organization.Effective_date = formatDateString(ba_organization.Effective_date)
	ba_organization.Expiry_date = formatDateString(ba_organization.Expiry_date)
	ba_organization.Removed_date = formatDateString(ba_organization.Removed_date)
	ba_organization.Row_effective_date = formatDateString(ba_organization.Row_effective_date)
	ba_organization.Row_expiry_date = formatDateString(ba_organization.Row_expiry_date)






	rows, err := stmt.Exec(ba_organization.Business_associate_id, ba_organization.Organization_id, ba_organization.Organization_seq_no, ba_organization.Active_ind, ba_organization.Address_obs_no, ba_organization.Address_source, ba_organization.Area_id, ba_organization.Area_type, ba_organization.Created_date, ba_organization.Description, ba_organization.Effective_date, ba_organization.Expiry_date, ba_organization.Organization_name, ba_organization.Organization_type, ba_organization.Ppdm_guid, ba_organization.Remark, ba_organization.Removed_date, ba_organization.Source, ba_organization.Row_changed_by, ba_organization.Row_changed_date, ba_organization.Row_created_by, ba_organization.Row_created_date, ba_organization.Row_effective_date, ba_organization.Row_expiry_date, ba_organization.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateBaOrganization(c *fiber.Ctx) error {
	var ba_organization dto.Ba_organization

	setDefaults(&ba_organization)

	if err := json.Unmarshal(c.Body(), &ba_organization); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ba_organization.Business_associate_id = id

    if ba_organization.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ba_organization set organization_id = :1, organization_seq_no = :2, active_ind = :3, address_obs_no = :4, address_source = :5, area_id = :6, area_type = :7, created_date = :8, description = :9, effective_date = :10, expiry_date = :11, organization_name = :12, organization_type = :13, ppdm_guid = :14, remark = :15, removed_date = :16, source = :17, row_changed_by = :18, row_changed_date = :19, row_created_by = :20, row_effective_date = :21, row_expiry_date = :22, row_quality = :23 where business_associate_id = :25")
	if err != nil {
		return err
	}

	ba_organization.Row_changed_date = formatDate(ba_organization.Row_changed_date)
	ba_organization.Created_date = formatDateString(ba_organization.Created_date)
	ba_organization.Effective_date = formatDateString(ba_organization.Effective_date)
	ba_organization.Expiry_date = formatDateString(ba_organization.Expiry_date)
	ba_organization.Removed_date = formatDateString(ba_organization.Removed_date)
	ba_organization.Row_effective_date = formatDateString(ba_organization.Row_effective_date)
	ba_organization.Row_expiry_date = formatDateString(ba_organization.Row_expiry_date)






	rows, err := stmt.Exec(ba_organization.Organization_id, ba_organization.Organization_seq_no, ba_organization.Active_ind, ba_organization.Address_obs_no, ba_organization.Address_source, ba_organization.Area_id, ba_organization.Area_type, ba_organization.Created_date, ba_organization.Description, ba_organization.Effective_date, ba_organization.Expiry_date, ba_organization.Organization_name, ba_organization.Organization_type, ba_organization.Ppdm_guid, ba_organization.Remark, ba_organization.Removed_date, ba_organization.Source, ba_organization.Row_changed_by, ba_organization.Row_changed_date, ba_organization.Row_created_by, ba_organization.Row_effective_date, ba_organization.Row_expiry_date, ba_organization.Row_quality, ba_organization.Business_associate_id)
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

func PatchBaOrganization(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ba_organization set "
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
		} else if key == "created_date" || key == "effective_date" || key == "expiry_date" || key == "removed_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where business_associate_id = :id"

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

func DeleteBaOrganization(c *fiber.Ctx) error {
	id := c.Params("id")
	var ba_organization dto.Ba_organization
	ba_organization.Business_associate_id = id

	stmt, err := db.Prepare("delete from ba_organization where business_associate_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ba_organization.Business_associate_id)
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


