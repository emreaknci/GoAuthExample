import React, { useContext } from 'react'
import { AuthContext } from '../contexts/AuthContext'
import { Typography } from '@mui/material'
import { JwtHelper } from '../helpers/jwtHelper'
import StorageService from '../services/storage.service'

const Home = () => {
  const { isAuthenticated, isTokenChecked } = useContext(AuthContext)
  return (
    <>
      <Typography variant="h2" component="h1" gutterBottom>
        Home
      </Typography>
      <Typography variant="h5" gutterBottom>
        {isTokenChecked ? 'Token is checked and' : 'Token is not checked and'} {isAuthenticated ? 'you are authenticated' : 'you are not authenticated'}
      </Typography>

      <Typography variant="h6" gutterBottom>
        {isAuthenticated && <>UserID: {JwtHelper.getUserId(StorageService.getAccessToken()!)}</>}
      </Typography>
    </>

  )
}

export default Home