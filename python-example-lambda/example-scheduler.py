def lambda_handler(event, context):
    print('Scheduled task executed')
    return {
        'statusCode': 200,
        'body': 'Task executed!'
    }
