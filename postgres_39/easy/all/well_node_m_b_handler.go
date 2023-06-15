package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellNodeMB(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_node_m_b")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_node_m_b

	for rows.Next() {
		var well_node_m_b dto.Well_node_m_b
		if err := rows.Scan(&well_node_m_b.Node_id, &well_node_m_b.Source, &well_node_m_b.Active_ind, &well_node_m_b.Dls_road_allowance_id, &well_node_m_b.Effective_date, &well_node_m_b.Ew_direction, &well_node_m_b.Ew_distance, &well_node_m_b.Ew_distance_ouom, &well_node_m_b.Ew_start_line, &well_node_m_b.Expiry_date, &well_node_m_b.Location_type, &well_node_m_b.Ns_direction, &well_node_m_b.Ns_distance, &well_node_m_b.Ns_distance_ouom, &well_node_m_b.Ns_start_line, &well_node_m_b.Orientation, &well_node_m_b.Parcel_carter_id, &well_node_m_b.Parcel_congress_id, &well_node_m_b.Parcel_dls_id, &well_node_m_b.Parcel_fps_id, &well_node_m_b.Parcel_libya_id, &well_node_m_b.Parcel_ne_loc_id, &well_node_m_b.Parcel_north_sea_id, &well_node_m_b.Parcel_nts_id, &well_node_m_b.Parcel_offshore_id, &well_node_m_b.Parcel_ohio_id, &well_node_m_b.Parcel_pbl_id, &well_node_m_b.Parcel_texas, &well_node_m_b.Ppdm_guid, &well_node_m_b.Reference_loc, &well_node_m_b.Remark, &well_node_m_b.Surface_loc, &well_node_m_b.Row_changed_by, &well_node_m_b.Row_changed_date, &well_node_m_b.Row_created_by, &well_node_m_b.Row_created_date, &well_node_m_b.Row_effective_date, &well_node_m_b.Row_expiry_date, &well_node_m_b.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_node_m_b to result
		result = append(result, well_node_m_b)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellNodeMB(c *fiber.Ctx) error {
	var well_node_m_b dto.Well_node_m_b

	setDefaults(&well_node_m_b)

	if err := json.Unmarshal(c.Body(), &well_node_m_b); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_node_m_b values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39)")
	if err != nil {
		return err
	}
	well_node_m_b.Row_created_date = formatDate(well_node_m_b.Row_created_date)
	well_node_m_b.Row_changed_date = formatDate(well_node_m_b.Row_changed_date)
	well_node_m_b.Effective_date = formatDateString(well_node_m_b.Effective_date)
	well_node_m_b.Expiry_date = formatDateString(well_node_m_b.Expiry_date)
	well_node_m_b.Row_effective_date = formatDateString(well_node_m_b.Row_effective_date)
	well_node_m_b.Row_expiry_date = formatDateString(well_node_m_b.Row_expiry_date)






	rows, err := stmt.Exec(well_node_m_b.Node_id, well_node_m_b.Source, well_node_m_b.Active_ind, well_node_m_b.Dls_road_allowance_id, well_node_m_b.Effective_date, well_node_m_b.Ew_direction, well_node_m_b.Ew_distance, well_node_m_b.Ew_distance_ouom, well_node_m_b.Ew_start_line, well_node_m_b.Expiry_date, well_node_m_b.Location_type, well_node_m_b.Ns_direction, well_node_m_b.Ns_distance, well_node_m_b.Ns_distance_ouom, well_node_m_b.Ns_start_line, well_node_m_b.Orientation, well_node_m_b.Parcel_carter_id, well_node_m_b.Parcel_congress_id, well_node_m_b.Parcel_dls_id, well_node_m_b.Parcel_fps_id, well_node_m_b.Parcel_libya_id, well_node_m_b.Parcel_ne_loc_id, well_node_m_b.Parcel_north_sea_id, well_node_m_b.Parcel_nts_id, well_node_m_b.Parcel_offshore_id, well_node_m_b.Parcel_ohio_id, well_node_m_b.Parcel_pbl_id, well_node_m_b.Parcel_texas, well_node_m_b.Ppdm_guid, well_node_m_b.Reference_loc, well_node_m_b.Remark, well_node_m_b.Surface_loc, well_node_m_b.Row_changed_by, well_node_m_b.Row_changed_date, well_node_m_b.Row_created_by, well_node_m_b.Row_created_date, well_node_m_b.Row_effective_date, well_node_m_b.Row_expiry_date, well_node_m_b.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellNodeMB(c *fiber.Ctx) error {
	var well_node_m_b dto.Well_node_m_b

	setDefaults(&well_node_m_b)

	if err := json.Unmarshal(c.Body(), &well_node_m_b); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_node_m_b.Node_id = id

    if well_node_m_b.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_node_m_b set source = :1, active_ind = :2, dls_road_allowance_id = :3, effective_date = :4, ew_direction = :5, ew_distance = :6, ew_distance_ouom = :7, ew_start_line = :8, expiry_date = :9, location_type = :10, ns_direction = :11, ns_distance = :12, ns_distance_ouom = :13, ns_start_line = :14, orientation = :15, parcel_carter_id = :16, parcel_congress_id = :17, parcel_dls_id = :18, parcel_fps_id = :19, parcel_libya_id = :20, parcel_ne_loc_id = :21, parcel_north_sea_id = :22, parcel_nts_id = :23, parcel_offshore_id = :24, parcel_ohio_id = :25, parcel_pbl_id = :26, parcel_texas = :27, ppdm_guid = :28, reference_loc = :29, remark = :30, surface_loc = :31, row_changed_by = :32, row_changed_date = :33, row_created_by = :34, row_effective_date = :35, row_expiry_date = :36, row_quality = :37 where node_id = :39")
	if err != nil {
		return err
	}

	well_node_m_b.Row_changed_date = formatDate(well_node_m_b.Row_changed_date)
	well_node_m_b.Effective_date = formatDateString(well_node_m_b.Effective_date)
	well_node_m_b.Expiry_date = formatDateString(well_node_m_b.Expiry_date)
	well_node_m_b.Row_effective_date = formatDateString(well_node_m_b.Row_effective_date)
	well_node_m_b.Row_expiry_date = formatDateString(well_node_m_b.Row_expiry_date)






	rows, err := stmt.Exec(well_node_m_b.Source, well_node_m_b.Active_ind, well_node_m_b.Dls_road_allowance_id, well_node_m_b.Effective_date, well_node_m_b.Ew_direction, well_node_m_b.Ew_distance, well_node_m_b.Ew_distance_ouom, well_node_m_b.Ew_start_line, well_node_m_b.Expiry_date, well_node_m_b.Location_type, well_node_m_b.Ns_direction, well_node_m_b.Ns_distance, well_node_m_b.Ns_distance_ouom, well_node_m_b.Ns_start_line, well_node_m_b.Orientation, well_node_m_b.Parcel_carter_id, well_node_m_b.Parcel_congress_id, well_node_m_b.Parcel_dls_id, well_node_m_b.Parcel_fps_id, well_node_m_b.Parcel_libya_id, well_node_m_b.Parcel_ne_loc_id, well_node_m_b.Parcel_north_sea_id, well_node_m_b.Parcel_nts_id, well_node_m_b.Parcel_offshore_id, well_node_m_b.Parcel_ohio_id, well_node_m_b.Parcel_pbl_id, well_node_m_b.Parcel_texas, well_node_m_b.Ppdm_guid, well_node_m_b.Reference_loc, well_node_m_b.Remark, well_node_m_b.Surface_loc, well_node_m_b.Row_changed_by, well_node_m_b.Row_changed_date, well_node_m_b.Row_created_by, well_node_m_b.Row_effective_date, well_node_m_b.Row_expiry_date, well_node_m_b.Row_quality, well_node_m_b.Node_id)
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

func PatchWellNodeMB(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_node_m_b set "
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
	query += " where node_id = :id"

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

func DeleteWellNodeMB(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_node_m_b dto.Well_node_m_b
	well_node_m_b.Node_id = id

	stmt, err := db.Prepare("delete from well_node_m_b where node_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_node_m_b.Node_id)
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


