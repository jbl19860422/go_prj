cmake_minimum_required(VERSION 2.8)
project(GOPRJ)

file(GLOB_RECURSE GODIRS 
    LIST_DIRECTORIES true
    ${CMAKE_SOURCE_DIR}/src/*)

foreach(GODIR in ${GODIRS})
    if(IS_DIRECTORY ${GODIR})
        message("dir=" ${GODIR})
        string(REGEX REPLACE "/" "_" GODIR_A ${GODIR})
        message("add target" ${GODIR_A})
        add_custom_target(${GODIR_A} ALL
            COMMAND go install 
            WORKING_DIRECTORY ${GODIR} 
        )
    endif()
endforeach()