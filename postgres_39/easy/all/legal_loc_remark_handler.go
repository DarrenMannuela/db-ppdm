package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetLegalLocRemark(c *fiber.Ctx) error {
	rows, err := db.Query("select * from legal_loc_remark")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Legal_loc_remark

	for rows.Next() {
		var legal_loc_remark dto.Legal_loc_remark
		if err := rows.Scan(&legal_loc_remark.Legal_loc_remark_id, &legal_loc_remark.Location_type, &legal_loc_remark.Source, &legal_loc_remark.Remark_seq_no, &legal_loc_remark.Active_ind, &legal_loc_remark.Effective_date, &legal_loc_remark.Expiry_date, &legal_loc_remark.Land_parcel_type, &legal_loc_remark.Legal_carter_loc_id, &legal_loc_remark.Legal_congress_loc_id, &legal_loc_remark.Legal_dls_loc_id, &legal_loc_remark.Legal_fps_loc_id, &legal_loc_remark.Legal_geodetic_loc_id, &legal_loc_remark.Legal_ne_loc_id, &legal_loc_remark.Legal_north_sea_loc_id, &legal_loc_remark.Legal_nts_loc_id, &legal_loc_remark.Legal_offshore_loc_id, &legal_loc_remark.Legal_ohio_loc_id, &legal_loc_remark.Legal_texas_loc_id, &legal_loc_remark.Ppdm_guid, &legal_loc_remark.Remark, &legal_loc_remark.Remark_type, &legal_loc_remark.Well_node_id, &legal_loc_remark.Row_changed_by, &legal_loc_remark.Row_changed_date, &legal_loc_remark.Row_created_by, &legal_loc_remark.Row_created_date, &legal_loc_remark.Row_effective_date, &legal_loc_remark.Row_expiry_date, &legal_loc_remark.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append legal_loc_remark to result
		result = append(result, legal_loc_remark)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetLegalLocRemark(c *fiber.Ctx) error {
	var legal_loc_remark dto.Legal_loc_remark

	setDefaults(&legal_loc_remark)

	if err := json.Unmarshal(c.Body(), &legal_loc_remark); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into legal_loc_remark values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30)")
	if err != nil {
		return err
	}
	legal_loc_remark.Row_created_date = formatDate(legal_loc_remark.Row_created_date)
	legal_loc_remark.Row_changed_date = formatDate(legal_loc_remark.Row_changed_date)
	legal_loc_remark.Effective_date = formatDateString(legal_loc_remark.Effective_date)
	legal_loc_remark.Expiry_date = formatDateString(legal_loc_remark.Expiry_date)
	legal_loc_remark.Row_effective_date = formatDateString(legal_loc_remark.Row_effective_date)
	legal_loc_remark.Row_expiry_date = formatDateString(legal_loc_remark.Row_expiry_date)






	rows, err := stmt.Exec(legal_loc_remark.Legal_loc_remark_id, legal_loc_remark.Location_type, legal_loc_remark.Source, legal_loc_remark.Remark_seq_no, legal_loc_remark.Active_ind, legal_loc_remark.Effective_date, legal_loc_remark.Expiry_date, legal_loc_remark.Land_parcel_type, legal_loc_remark.Legal_carter_loc_id, legal_loc_remark.Legal_congress_loc_id, legal_loc_remark.Legal_dls_loc_id, legal_loc_remark.Legal_fps_loc_id, legal_loc_remark.Legal_geodetic_loc_id, legal_loc_remark.Legal_ne_loc_id, legal_loc_remark.Legal_north_sea_loc_id, legal_loc_remark.Legal_nts_loc_id, legal_loc_remark.Legal_offshore_loc_id, legal_loc_remark.Legal_ohio_loc_id, legal_loc_remark.Legal_texas_loc_id, legal_loc_remark.Ppdm_guid, legal_loc_remark.Remark, legal_loc_remark.Remark_type, legal_loc_remark.Well_node_id, legal_loc_remark.Row_changed_by, legal_loc_remark.Row_changed_date, legal_loc_remark.Row_created_by, legal_loc_remark.Row_created_date, legal_loc_remark.Row_effective_date, legal_loc_remark.Row_expiry_date, legal_loc_remark.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateLegalLocRemark(c *fiber.Ctx) error {
	var legal_loc_remark dto.Legal_loc_remark

	setDefaults(&legal_loc_remark)

	if err := json.Unmarshal(c.Body(), &legal_loc_remark); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	legal_loc_remark.Legal_loc_remark_id = id

    if legal_loc_remark.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update legal_loc_remark set location_type = :1, source = :2, remark_seq_no = :3, active_ind = :4, effective_date = :5, expiry_date = :6, land_parcel_type = :7, legal_carter_loc_id = :8, legal_congress_loc_id = :9, legal_dls_loc_id = :10, legal_fps_loc_id = :11, legal_geodetic_loc_id = :12, legal_ne_loc_id = :13, legal_north_sea_loc_id = :14, legal_nts_loc_id = :15, legal_offshore_loc_id = :16, legal_ohio_loc_id = :17, legal_texas_loc_id = :18, ppdm_guid = :19, remark = :20, remark_type = :21, well_node_id = :22, row_changed_by = :23, row_changed_date = :24, row_created_by = :25, row_effective_date = :26, row_expiry_date = :27, row_quality = :28 where legal_loc_remark_id = :30")
	if err != nil {
		return err
	}

	legal_loc_remark.Row_changed_date = formatDate(legal_loc_remark.Row_changed_date)
	legal_loc_remark.Effective_date = formatDateString(legal_loc_remark.Effective_date)
	legal_loc_remark.Expiry_date = formatDateString(legal_loc_remark.Expiry_date)
	legal_loc_remark.Row_effective_date = formatDateString(legal_loc_remark.Row_effective_date)
	legal_loc_remark.Row_expiry_date = formatDateString(legal_loc_remark.Row_expiry_date)






	rows, err := stmt.Exec(legal_loc_remark.Location_type, legal_loc_remark.Source, legal_loc_remark.Remark_seq_no, legal_loc_remark.Active_ind, legal_loc_remark.Effective_date, legal_loc_remark.Expiry_date, legal_loc_remark.Land_parcel_type, legal_loc_remark.Legal_carter_loc_id, legal_loc_remark.Legal_congress_loc_id, legal_loc_remark.Legal_dls_loc_id, legal_loc_remark.Legal_fps_loc_id, legal_loc_remark.Legal_geodetic_loc_id, legal_loc_remark.Legal_ne_loc_id, legal_loc_remark.Legal_north_sea_loc_id, legal_loc_remark.Legal_nts_loc_id, legal_loc_remark.Legal_offshore_loc_id, legal_loc_remark.Legal_ohio_loc_id, legal_loc_remark.Legal_texas_loc_id, legal_loc_remark.Ppdm_guid, legal_loc_remark.Remark, legal_loc_remark.Remark_type, legal_loc_remark.Well_node_id, legal_loc_remark.Row_changed_by, legal_loc_remark.Row_changed_date, legal_loc_remark.Row_created_by, legal_loc_remark.Row_effective_date, legal_loc_remark.Row_expiry_date, legal_loc_remark.Row_quality, legal_loc_remark.Legal_loc_remark_id)
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

func PatchLegalLocRemark(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update legal_loc_remark set "
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
	query += " where legal_loc_remark_id = :id"

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

func DeleteLegalLocRemark(c *fiber.Ctx) error {
	id := c.Params("id")
	var legal_loc_remark dto.Legal_loc_remark
	legal_loc_remark.Legal_loc_remark_id = id

	stmt, err := db.Prepare("delete from legal_loc_remark where legal_loc_remark_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(legal_loc_remark.Legal_loc_remark_id)
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


