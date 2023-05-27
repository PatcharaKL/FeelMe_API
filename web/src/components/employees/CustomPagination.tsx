import React, { useState } from "react";
import { Pagination } from "@mui/material";

  const CustomPagination = ({ itemsPerPage, totalItems, paginate }:any) => {
    const [currentPage, setCurrentPage] = useState(1);
    const handlePageChange = (event:any, page:any) => {
      setCurrentPage(page);
      paginate(page);
    };
  
    return (
      <Pagination
        className="flex justify-end"
        color="secondary"
        shape="rounded"
        count={Math.ceil(totalItems / itemsPerPage)}
        page={currentPage}
        onChange={handlePageChange}
      />
    );
  };

export default CustomPagination;
